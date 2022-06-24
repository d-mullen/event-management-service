package rowkey

import (
	"bytes"
	"sync"
	"time"

	"io"

	"github.com/pkg/errors"
	zbytes "github.com/zenoss/zingo/v4/bytes"
	"github.com/zenoss/zingo/v4/orderedbytes"
)

// ErrUnsupportedDecodeType represents a failure to decode because we
// haven't supplied a decoding function for the type requested.
var (
	ErrUnsupportedDecodeType = errors.New("cannot decode into provided type")
	ErrInvalidOffset         = errors.New("invalid offset")
)

type RowKeyDecoder interface {
	// Read attempts to read the value at the current offset and
	// decode it, replacing the value at dst, which should be
	// a pointer to an instance of a supported type. If the type of dst is not
	// supported, ErrUnsupportedDecodeType is returned. If there's no value at
	// the current offset, io.EOF is returned
	Read(dst interface{}) error

	// Get gets the encoded value at the current offset. A slice backed by the
	// original array is returned. If there's no value at the current offset,
	// io.EOF is returned.
	Get() ([]byte, error)

	// GetN returns a slice containing the next n values, or io.EOF if fewer than n
	// fields remain. The slice is backed by the original array.
	GetN(n int) ([]byte, error)

	// Match checks if the byte slice provided matches the buffer,
	// starting from the current offset. The provided byte slice must represent
	// complete, orderedbytes-encoded values.
	Match(cmp []byte) bool

	// ReadUntyped attempts to read the value at the current offset, decode it,
	// and return it.  If the type of dst is not, ErrUnsupportedDecodeType is
	// returned. If there's no value at the current offset, io.EOF is returned.
	ReadUntyped() (interface{}, error)

	// Seek implements io.Seek
	// Seek sets the offset for the next Read to offset,
	// interpreted according to whence:
	//   SeekStart means relative to the start of the file,
	//   SeekCurrent means relative to the current offset, and
	//   SeekEnd means relative to the end.
	// Seek returns the new offset relative to the start of the rowkey and an error, if any.
	// Seeking to an offset before the start of the rowkey is an error. Seeking to any positive offset is legal,
	// but subsequent reads will return io.EOF if the offset is beyond the end of the rowkey
	Seek(offset int64, whence int) (int64, error)

	// Len returns the number of bytes of the unread portion of the buffer
	Len() int

	// Free releases the decoder to be used again
	Free()
}

// NewRowKeyDecoder returns the default implementation of RowKeyDecoder, initialized to decode the rowkey in b
func NewRowKeyDecoder(b []byte) RowKeyDecoder {
	d := rkdpool.Get().(*rowKeyDecoder)
	d.buff = b
	return d
}

func NewRowKeyDecoderFromString(s string) RowKeyDecoder {
	// Convert directly from string to byte slice without a copy. This is safe
	// because we never modify the byte slice in the decoder.
	return NewRowKeyDecoder(zbytes.StringToBytesUnsafe(s))
}

var rkdpool = sync.Pool{
	New: func() interface{} {
		offsets := make([]int, 1, 10)
		offsets[0] = 0
		return &rowKeyDecoder{
			offsets:            offsets,
			currentFieldOffset: 0,
		}
	},
}

type rowKeyDecoder struct {
	buff               []byte
	offsets            []int
	currentFieldOffset int64
	currentSliceIndex  int
}

// if this is true then the last index stored in offsets will point to the end of the underlying slice
func (d *rowKeyDecoder) foundEnd() bool {
	if len(d.offsets) > 0 {
		return d.offsets[len(d.offsets)-1] >= len(d.buff)
	}
	return false
}

func (d *rowKeyDecoder) Free() {
	d.buff = nil
	d.offsets = d.offsets[:1]
	d.currentFieldOffset = 0
	d.currentSliceIndex = 0
	rkdpool.Put(d)
}

func (d *rowKeyDecoder) Read(dst interface{}) error {
	err := d.decode(dst)
	if err != nil {
		return errors.WithStack(err)
	}

	d.incrementCurrentPosition()

	return nil
}

func (d *rowKeyDecoder) Get() ([]byte, error) {
	return d.GetN(1)
}

func (d *rowKeyDecoder) GetN(n int) (slice []byte, err error) {
	var (
		i, next, size int
		start         = d.currentSliceIndex
	)
	for i = 0; i < n; i++ {
		if d.currentSliceIndex >= len(d.buff) {
			return nil, io.EOF
		}
		size, err = orderedbytes.Size(d.buff[d.currentSliceIndex:])
		if err != nil {
			return nil, errors.Wrap(err, "error finding size of next value")
		}
		next = d.currentSliceIndex + size
		d.currentSliceIndex = next
		d.incrementCurrentPosition()
	}
	slice = d.buff[start:next]
	return
}

func (d *rowKeyDecoder) Match(cmp []byte) bool {
	return bytes.HasPrefix(d.buff[d.currentSliceIndex:], cmp)
}

func (d *rowKeyDecoder) ReadUntyped() (interface{}, error) {
	if d.currentSliceIndex >= len(d.buff) {
		return nil, io.EOF
	}
	i, c, err := orderedbytes.DecodeFromSlice(d.buff[d.currentSliceIndex:])
	if err != nil {
		return nil, errors.Wrap(err, "error decoding next value")
	}
	d.currentSliceIndex += c

	// Update the offsets list
	d.incrementCurrentPosition()

	return i, nil
}

func (d *rowKeyDecoder) Seek(offset int64, whence int) (int64, error) {
	// Compute the desired offset from start
	var offsetFromStart int64
	switch whence {
	case io.SeekStart:
		offsetFromStart = offset
	case io.SeekCurrent:
		offsetFromStart = d.currentFieldOffset + offset
	case io.SeekEnd:
		if !d.foundEnd() {
			// Seek to the end so we know where it is
			err := d.seekEnd()
			if err != nil {
				return 0, errors.Wrap(err, "error seeking end of rowkey")
			}
		}
		// When we know the end, the last entry in d.offsets will be the end of the underlying buffer
		offsetFromStart = int64(len(d.offsets)-1) + offset
	default:
		return 0, errors.Wrap(ErrInvalidOffset, "uknown whence value")
	}

	// A negative offset-from-start is an error
	if offsetFromStart < 0 {
		return 0, errors.Wrap(ErrInvalidOffset, "offset is before start of rowkey")
	}

	// If we are already at the correct offset, there's nothing to do
	if offsetFromStart == d.currentFieldOffset {
		return d.currentFieldOffset, nil
	}

	// Jump ahead to the nearest offset we know about
	offsetToSeek := offsetFromStart
	if int64(len(d.offsets)-1) < offsetToSeek {
		offsetToSeek = int64(len(d.offsets) - 1)
	}

	if offsetToSeek >= 0 {
		d.currentSliceIndex = d.offsets[offsetToSeek]
		d.currentFieldOffset = offsetToSeek
	}

	// Seek forward the slow way to get to our desired offset
	if d.currentFieldOffset < offsetFromStart {
		err := d.seekForward(offsetFromStart - d.currentFieldOffset)
		if errors.Cause(err) == io.EOF {
			// In the case of an EOF, just return the end offset, which is where we currently are
			return d.currentFieldOffset, nil
		} else if err != nil {
			return 0, errors.WithStack(err)
		}
	}

	return offsetFromStart, nil
}

func (d *rowKeyDecoder) Len() int {
	return len(d.buff) - d.currentSliceIndex
}

func (d *rowKeyDecoder) seekForward(offset int64) error {
	for i := int64(0); i < offset; i++ {
		// Decode an object to move the cursor forward
		_, err := d.ReadUntyped()
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (d *rowKeyDecoder) incrementCurrentPosition() {
	// increment the current offset and store our current slice index at this offset
	d.currentFieldOffset++
	if int64(len(d.offsets)) < d.currentFieldOffset+1 {
		// update our offsets list
		d.offsets = append(d.offsets, d.currentSliceIndex)
	}
}

func (d *rowKeyDecoder) seekEnd() error {
	// Seek forward until we hit EOF
	var err error
	for !d.foundEnd() {
		err = d.seekForward(1)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (d *rowKeyDecoder) decode(dst interface{}) error {
	switch dst := dst.(type) {
	case *int:
		i, c, e := orderedbytes.DecodeInt32FromSlice(d.buff[d.currentSliceIndex:])
		if e != nil {
			return errors.WithStack(e)
		}
		d.currentSliceIndex += c
		*dst = int(i)
	case *int32:
		i, c, e := orderedbytes.DecodeInt32FromSlice(d.buff[d.currentSliceIndex:])
		if e != nil {
			return errors.WithStack(e)
		}
		d.currentSliceIndex += c
		*dst = i
	case *int64:
		i, c, e := orderedbytes.DecodeInt64FromSlice(d.buff[d.currentSliceIndex:])
		if e != nil {
			return errors.WithStack(e)
		}
		d.currentSliceIndex += c
		*dst = i
	case *float64:
		i, c, e := orderedbytes.DecodeFloat64FromSlice(d.buff[d.currentSliceIndex:])
		if e != nil {
			return errors.WithStack(e)
		}
		d.currentSliceIndex += c
		*dst = i
	case *float32:
		i, c, e := orderedbytes.DecodeFloat32FromSlice(d.buff[d.currentSliceIndex:])
		if e != nil {
			return errors.WithStack(e)
		}
		d.currentSliceIndex += c
		*dst = i
	case *bool:
		intval, c, e := orderedbytes.DecodeInt8FromSlice(d.buff[d.currentSliceIndex:])
		if e != nil {
			return errors.Wrap(e, "error decoding int8 value for bool")
		}
		d.currentSliceIndex += c
		*dst = intval == 1
	case *string:
		i, c, e := orderedbytes.DecodeStringFromSlice(d.buff[d.currentSliceIndex:])
		if e != nil {
			return errors.WithStack(e)
		}
		d.currentSliceIndex += c
		*dst = i
	case *time.Time:
		t, c, e := orderedbytes.DecodeInt64FromSlice(d.buff[d.currentSliceIndex:])
		if e != nil {
			return errors.WithStack(e)
		}
		d.currentSliceIndex += c
		*dst = time.Unix(0, t*int64(time.Millisecond))
	default:
		return ErrUnsupportedDecodeType
	}

	return nil
}
