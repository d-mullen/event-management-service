// Package rowkey provides utilities for encoding values to and decoding values
// from delimited strings, intended for creating and parsing BigTable row keys.
// Regardless of the encoding used for a particular type, lexicographical order
// is preserved.
package rowkey

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/zenoss/zingo/v4/orderedbytes"
)

// ErrUnsupportedEncodeType represents a failure to encode because we
// haven't supplied an encoding function for the type requested.
var ErrUnsupportedEncodeType = errors.New("cannot encode provided type")

// Buffer is an interface implemented by bytes.Buffer. It is used to back
// RowKeyEncoder so we can swap in a mock buffer that throws write errors in
// tests.
type Buffer interface {
	io.Writer
	fmt.Stringer
	Reset()
	Bytes() []byte
	Len() int
}

// RowKeyEncoder is used to build a row key containing delimited,
// order-preserving-encoded values. It is not safe for concurrent use, but why
// should it be?
type RowKeyEncoder interface {

	// WithBuffer sets the buffer backing the encoder. It returns the encoder.
	// If not called, a bytes.Buffer is used. Note: a call to WithBuffer will
	// reset the encoder.
	WithBuffer(Buffer) RowKeyEncoder

	// Write encodes and writes the provided value to the buffer. Supported
	// types are int, int32, int64, float64, string, and time.Time.
	// Invoking Write with an unsupported type will result in an error.
	Write(interface{}) (int, error)

	// MustWrite is like Write, but it panics on error
	MustWrite(interface{})

	// WriteRaw writes the raw bytes directly to the buffer with no additional encoding
	// or validation.  Use at your own risk.
	WriteRaw([]byte) (int, error)

	// MustWriteRaw is like WriteRaw, but it panics on error
	MustWriteRaw([]byte)

	// WriteReverseOrder performs a Write, but encodes the byte array in
	// a manner that results in lexicographical order being
	// reversed.
	WriteReverseOrder(interface{}) (int, error)

	// MustWriteReverseOrder is like WriteReverseOrder, but it panics on error
	MustWriteReverseOrder(interface{})

	// WriteMaxLen encodes and writes the provided value to the buffer, but limits
	// the actual bytes written to a maximum length provided.
	WriteMaxLen(interface{}, int) (int, error)

	// MustWriteMaxLen is like WriteMaxLen, but it panics on error
	MustWriteMaxLen(interface{}, int)

	// Bytes returns the slice of the underlying Buffer containing the key in
	// its current state. Note: the default implementation does not return
	// a copy of the byte slice, but the slice itself. Reusing or resetting the
	// same encoder will result on the value returned by this call being
	// altered.
	Bytes() []byte

	// Len returns the length of the encoded key.
	Len() int

	// Reset resets the encoder to a clean state, emptying the buffer
	Reset()

	fmt.Stringer

	// BytesPrefix returns the byte slice representation of the key. It differs
	// from Bytes in that it strips the terminator byte introduced by the
	// encoding, if present.
	BytesPrefix() []byte

	// StringPrefix returns the string representation of the key. It differs
	// from String in that it strips the terminator byte introduced by the
	// encoding, if present.
	StringPrefix() string
}

// NewRowKeyEncoder returns an instance of the default implementation of
// RowKeyEncoder, initialized with a bytes.Buffer as the buffer.
func NewRowKeyEncoder() RowKeyEncoder {
	var buf bytes.Buffer
	return &rowKeyEncoder{buf: &buf}
}

// NewRowKeyEncoderFromString returns an instance of the default implementation of
// RowKeyEncoder, initialized with rowkey already in the underlying buffer.
func NewRowKeyEncoderFromString(rowkey string) RowKeyEncoder {
	buf := bytes.NewBufferString(rowkey)
	return &rowKeyEncoder{buf: buf}
}

func NewRowKeyEncoderWithCapacity(cap int) RowKeyEncoder {
	slice := make([]byte, 0, cap)
	buf := bytes.NewBuffer(slice)
	return &rowKeyEncoder{buf: buf}
}

type rowKeyEncoder struct {
	buf        Buffer
	terminated bool
}

func (r *rowKeyEncoder) WithBuffer(b Buffer) RowKeyEncoder {
	r.Reset()
	r.buf = b
	return r
}

// Write writes bytes to the buffer.
func (r *rowKeyEncoder) Write(p interface{}) (int, error) {
	b, err := r.encode(p, orderedbytes.Ascending)
	if err != nil {
		return 0, err
	}
	return r.write(b, p)
}

func (r *rowKeyEncoder) MustWrite(p interface{}) {
	_, err := r.Write(p)
	if err != nil {
		panic(err)
	}
}

func (r *rowKeyEncoder) WriteRaw(b []byte) (int, error) {
	return r.buf.Write(b)
}

func (r *rowKeyEncoder) MustWriteRaw(b []byte) {
	_, err := r.WriteRaw(b)
	if err != nil {
		panic(err)
	}
}

func (r *rowKeyEncoder) WriteMaxLen(p interface{}, max int) (int, error) {
	// Strings have a terminator byte, so we'll need to truncate the value
	// itself, rather than the encoded value
	var isstr bool
	if s, ok := p.(string); ok && max > 2 && len(s)+2 > max {
		isstr = true
		p = s[:max-2]
	}
	b, err := r.encode(p, orderedbytes.Ascending)
	if err != nil {
		return 0, err
	}
	if !isstr && max > 0 && len(b) > max {
		b = b[:max]
	}
	return r.write(b, p)
}

func (r *rowKeyEncoder) MustWriteMaxLen(p interface{}, max int) {
	_, err := r.WriteMaxLen(p, max)
	if err != nil {
		panic(err)
	}
}

func (r *rowKeyEncoder) WriteReverseOrder(p interface{}) (int, error) {
	b, err := r.encode(p, orderedbytes.Descending)
	if err != nil {
		return 0, err
	}
	return r.write(b, p)
}

func (r *rowKeyEncoder) MustWriteReverseOrder(p interface{}) {
	_, err := r.WriteReverseOrder(p)
	if err != nil {
		panic(err)
	}
}

func (r *rowKeyEncoder) write(b []byte, orig interface{}) (int, error) {
	n, err := r.buf.Write(b)
	if err != nil {
		return 0, err
	}
	r.setTerminated(orig)
	return n, nil
}

func (r *rowKeyEncoder) encode(v interface{}, order orderedbytes.Order) ([]byte, error) {
	var b []byte
	switch v := v.(type) {
	case time.Time:
		b = orderedbytes.EncodeInt64(toMillis(v), order)
	case bool:
		val := int8(0)
		if v {
			val = int8(1)
		}
		b = orderedbytes.EncodeInt8(val, order)
	default:
		var err error
		b, err = orderedbytes.Encode(v, order)
		if err != nil {
			if err == orderedbytes.ErrUnsupportedType {
				return nil, ErrUnsupportedEncodeType
			}
			return nil, err
		}
	}
	return b, nil
}

func (r *rowKeyEncoder) String() string {
	return r.buf.String()
}

func (r *rowKeyEncoder) StringPrefix() string {
	return string(r.BytesPrefix())
}

func (r *rowKeyEncoder) Bytes() []byte {
	// We have to copy the slice, because the buffer will
	//  re-use the backing array for future resets/writes
	b := r.buf.Bytes()
	res := make([]byte, len(b))
	copy(res, b)
	return res
}

func (r *rowKeyEncoder) BytesPrefix() []byte {
	b := r.Bytes()
	if r.terminated {
		// Strip off the terminator byte
		b = b[:len(b)-1]
	}
	return b
}

func (r *rowKeyEncoder) Reset() {
	r.buf.Reset()
}

func (r *rowKeyEncoder) Len() int {
	return r.buf.Len()
}

func (r *rowKeyEncoder) setTerminated(p interface{}) {
	_, r.terminated = p.(string)
}

func toMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
