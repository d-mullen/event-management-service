# rowkey
--
    import "github.com/zenoss/zingo/v4/rowkey"

Package rowkey provides utilities for encoding values to and decoding values
from strings, intended for creating and parsing BigTable row keys.
Regardless of the encoding used for a particular type, lexicographical order is
preserved.

## Usage

```go
var ErrInvalidIndex = errors.New("index is inappropriate for this row key")
```
ErrInvalidIndex is returned when a call to RowKeyDecoder.Read is made with an
index that doesn't have a corresponding value in the key.

```go
var ErrUnsupportedDecodeType = errors.New("cannot decode into provided type")
```
ErrUnsupportedDecodeType represents a failure to decode because we haven't
supplied a decoding function for the type requested.

```go
var ErrUnsupportedEncodeType = errors.New("cannot encode provided type")
```
ErrUnsupportedEncodeType represents a failure to encode because we haven't
supplied an encoding function for the type requested.

#### type Buffer

```go
type Buffer interface {
	io.Writer
	fmt.Stringer
	Reset()
	Bytes() []byte
	Len() int
}
```

Buffer is an interface implemented by bytes.Buffer. It is used to back
RowKeyEncoder so we can swap in a mock buffer that throws write errors in tests.

#### type RowKeyDecoder

```go
type RowKeyDecoder interface {
    // Read attempts to read the value at the current offset and
    // decode it, replacing the value at dst, which should be
    // a pointer to an instance of a supported type. If the type of dst is not
    // supported, ErrUnsupportedDecodeType is returned. If there's no value at
    // the current offset, io.EOF is returned
    Read(dst interface{}) error

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
}
```


#### func  NewRowKeyDecoder

```go
func NewRowKeyDecoder(b []byte) RowKeyDecoder
```
NewRowKeyDecoder returns the default implementation of RowKeyDecoder, initialized to decode the rowkey in the input byte slice

#### type RowKeyEncoder

```go
type RowKeyEncoder interface {

	// WithBuffer sets the buffer backing the encoder. It returns the encoder.
    // If not called, a bytes.Buffer is used. Note: a call to WithBuffer will
    // reset the encoder.
    WithBuffer(Buffer) RowKeyEncoder

    // Write encodes and writes the provided value to the buffer. Supported 
    // types are int, int32, int64, float64, string, and time.Time. 
    // Invoking Write with an unsupported type will result in an error.
    Write(interface{}) (int, error)

    // WriteReverseOrder performs a Write, but encodes the byte array in 
    // a manner that results in lexicographical order being
    // reversed.
    WriteReverseOrder(interface{}) (int, error)

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
}
```

RowKeyEncoder is used to build a row key containing multiple
order-preserving-encoded values. It is not safe for concurrent use, but why
should it be?

#### func  NewRowKeyEncoder

```go
func NewRowKeyEncoder() RowKeyEncoder
```
NewRowKeyEncoder returns an instance of the default implementation of
RowKeyEncoder, initialized with a bytes.Buffer as the buffer.
