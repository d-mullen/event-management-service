// Encode various types such that the lexical ordering of the encoded bytes
// matchees the natural ordering of the original types.
//
// This encoding format is implemented as a port of org.apache.hadoop.hbase.util.OrderedBytes

package orderedbytes

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"strings"

	"github.com/pkg/errors"
	zbytes "github.com/zenoss/zingo/v4/bytes"
)

const (
	Null             byte = 0x05
	FixedInt8        byte = 0x29
	FixedInt16       byte = 0x2a
	FixedInt32       byte = 0x2b
	FixedInt64       byte = 0x2c
	FixedFloat32     byte = 0x30
	FixedFloat64     byte = 0x31
	Text             byte = 0x34
	StringTerminator byte = 0x00
)

type Order int8

const (
	Ascending  Order = 0
	Descending Order = 1
)

var (
	nullDesc             byte = xorbyte(Null)
	fixedInt8Desc        byte = xorbyte(FixedInt8)
	fixedInt16Desc       byte = xorbyte(FixedInt16)
	fixedInt32Desc       byte = xorbyte(FixedInt32)
	fixedInt64Desc       byte = xorbyte(FixedInt64)
	fixedFloat32Desc     byte = xorbyte(FixedFloat32)
	fixedFloat64Desc     byte = xorbyte(FixedFloat64)
	textDesc             byte = xorbyte(Text)
	stringTerminatorDesc byte = xorbyte(StringTerminator)

	ErrUnsupportedType          = errors.New("unsupported type")
	ErrWrongTypeCode            = errors.New("wrong type code")
	ErrStringContainsTerminator = errors.New("string contains null byte")
)

///////////////////////////////////////////////////////////////////////////////
// Encode an integer type such that its encoded bytewise ordering matches its
// natural ordering.
//
///////////////////////////////////////////////////////////////////////////////
func Encode(v interface{}, order Order) ([]byte, error) {
	var b []byte
	switch v := v.(type) {
	case int8:
		b = EncodeInt8(v, order)
	case int16:
		b = EncodeInt16(v, order)
	case int32:
		b = EncodeInt32(v, order)
	case int:
		b = EncodeInt32(int32(v), order)
	case uint32, uint64:
		// Unsigned integers would be easy to support, but the java package we are porting
		//  doesn't have an op code for them.
		return nil, ErrUnsupportedType
	case int64:
		b = EncodeInt64(v, order)
	case float64:
		b = EncodeFloat64(v, order)
	case float32:
		b = EncodeFloat32(v, order)
	case string:
		var err error
		b, err = EncodeString(v, order)
		if err != nil {
			return nil, errors.Wrap(err, "error encoding string value")
		}
	case nil:
		b = EncodeNull(order)
	default:
		return nil, ErrUnsupportedType
	}
	return b, nil
}

func Decode(buffer io.ReadSeeker) (interface{}, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return "", err
	}
	if count < 1 {
		return "", io.EOF
	}
	code := tmp[0]

	switch code {
	case FixedInt8:
		return decodeInt8(buffer, Ascending)
	case fixedInt8Desc:
		return decodeInt8(buffer, Descending)
	case FixedInt16:
		return decodeInt16(buffer, Ascending)
	case fixedInt16Desc:
		return decodeInt16(buffer, Descending)
	case FixedInt32:
		return decodeInt32(buffer, Ascending)
	case fixedInt32Desc:
		return decodeInt32(buffer, Descending)
	case FixedInt64:
		return decodeInt64(buffer, Ascending)
	case fixedInt64Desc:
		return decodeInt64(buffer, Descending)
	case FixedFloat64:
		return decodeFloat64(buffer, Ascending)
	case fixedFloat64Desc:
		return decodeFloat64(buffer, Descending)
	case FixedFloat32:
		return decodeFloat32(buffer, Ascending)
	case fixedFloat32Desc:
		return decodeFloat32(buffer, Descending)
	case Text:
		return decodeString(buffer, Ascending)
	case textDesc:
		return decodeString(buffer, Descending)
	case Null, nullDesc:
		return nil, nil
	}
	return nil, ErrUnsupportedType
}

func DecodeFromString(s string) (interface{}, int, error) {
	return DecodeFromSlice(zbytes.StringToBytesUnsafe(s))
}

func DecodeFromSlice(b []byte) (interface{}, int, error) {
	if len(b) < 1 {
		return "", 0, io.EOF
	}
	code := b[0]

	var (
		obj  interface{}
		read int
		err  error
	)

	switch code {
	case FixedInt8:
		obj, read, err = decodeInt8FromSlice(b[1:], Ascending)
	case fixedInt8Desc:
		obj, read, err = decodeInt8FromSlice(b[1:], Descending)
	case FixedInt16:
		obj, read, err = decodeInt16FromSlice(b[1:], Ascending)
	case fixedInt16Desc:
		obj, read, err = decodeInt16FromSlice(b[1:], Descending)
	case FixedInt32:
		obj, read, err = decodeInt32FromSlice(b[1:], Ascending)
	case fixedInt32Desc:
		obj, read, err = decodeInt32FromSlice(b[1:], Descending)
	case FixedInt64:
		obj, read, err = decodeInt64FromSlice(b[1:], Ascending)
	case fixedInt64Desc:
		obj, read, err = decodeInt64FromSlice(b[1:], Descending)
	case FixedFloat64:
		obj, read, err = decodeFloat64FromSlice(b[1:], Ascending)
	case fixedFloat64Desc:
		obj, read, err = decodeFloat64FromSlice(b[1:], Descending)
	case FixedFloat32:
		obj, read, err = decodeFloat32FromSlice(b[1:], Ascending)
	case fixedFloat32Desc:
		obj, read, err = decodeFloat32FromSlice(b[1:], Descending)
	case Text:
		obj, read, err = decodeStringFromSlice(b[1:], Ascending)
	case textDesc:
		obj, read, err = decodeStringFromSlice(b[1:], Descending)
	case Null, nullDesc:
		obj, read, err = nil, 0, nil
	default:
		return nil, 0, ErrUnsupportedType
	}
	return obj, read + 1, err
}

func EncodeNull(order Order) []byte {
	var tmp = make([]byte, 1)
	if order == Ascending {
		tmp[0] = Null
	} else {
		tmp[0] = nullDesc
	}
	return tmp
}

///////////////////////////////////////////////////////////////////////////////
// Encode an integer type such that its encoded bytewise ordering matches
// its natural ordering.
///////////////////////////////////////////////////////////////////////////////

// Encode an 8-bit integer
func EncodeInt8(val int8, order Order) []byte {
	var tmp = make([]byte, 2)
	tmp[0] = FixedInt8
	tmp[1] = byte(uint8(val) ^ 1<<7)
	if order == Descending {
		xorbytes(tmp)
	}
	return tmp
}

// Decode an 8-bit integer stored in the given buffer
func DecodeInt8(buffer io.ReadSeeker) (int8, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return int8(0), err
	}
	if count < 1 {
		return int8(0), io.EOF
	}
	code := tmp[0]
	if code != FixedInt8 && code != fixedInt8Desc {
		return int8(0), ErrWrongTypeCode
	}
	if code == fixedInt8Desc {
		return decodeInt8(buffer, Descending)
	}
	return decodeInt8(buffer, Ascending)
}

func decodeInt8(buffer io.ReadSeeker, order Order) (int8, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return int8(0), err
	}
	if count < 1 {
		return int8(0), io.EOF
	}

	val, _, err := decodeInt8FromSlice(tmp, order)
	return val, err
}

func DecodeInt8FromSlice(b []byte) (int8, int, error) {
	if len(b) < 1 {
		return int8(0), 0, io.EOF
	}
	code := b[0]
	if code != FixedInt8 && code != fixedInt8Desc {
		return int8(0), 0, ErrWrongTypeCode
	}
	order := Ascending
	if code == fixedInt8Desc {
		order = Descending
	}
	val, offset, err := decodeInt8FromSlice(b[1:], order)
	return val, offset + 1, err
}

func DecodeInt8FromString(s string) (int8, int, error) {
	return DecodeInt8FromSlice(zbytes.StringToBytesUnsafe(s))
}

func decodeInt8FromSlice(b []byte, order Order) (int8, int, error) {
	if len(b) < 1 {
		return int8(0), 0, io.EOF
	}

	bite := b[0]
	if order == Descending {
		bite = xorbyte(bite)
	}

	var val int8 = int8((bite ^ 0x80) & 0xff)
	return val, 1, nil
}

// Encode a 16-bit integer
func EncodeInt16(val int16, order Order) []byte {
	var tmp = make([]byte, 3)
	tmp[0] = FixedInt16
	binary.BigEndian.PutUint16(tmp[1:], uint16(val)^(1<<15))
	if order == Descending {
		xorbytes(tmp)
	}
	return tmp
}

// Decode a 16-bit integer stored in the given buffer
func DecodeInt16(buffer io.ReadSeeker) (int16, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return int16(0), err
	}
	if count < 1 {
		return int16(0), io.EOF
	}
	code := tmp[0]
	if code != FixedInt16 && code != fixedInt16Desc {
		return int16(0), ErrWrongTypeCode
	}
	if code == fixedInt16Desc {
		return decodeInt16(buffer, Descending)
	}
	return decodeInt16(buffer, Ascending)
}

func decodeInt16(buffer io.ReadSeeker, order Order) (int16, error) {
	tmp := make([]byte, 2)
	count, err := buffer.Read(tmp)
	if err != nil {
		return int16(0), err
	}
	if count < 2 {
		return int16(0), io.EOF
	}

	val, _, err := decodeInt16FromSlice(tmp, order)

	return val, err
}

func DecodeInt16FromSlice(b []byte) (int16, int, error) {
	if len(b) < 1 {
		return int16(0), 0, io.EOF
	}
	code := b[0]
	if code != FixedInt16 && code != fixedInt16Desc {
		return int16(0), 0, ErrWrongTypeCode
	}
	order := Ascending
	if code == fixedInt16Desc {
		order = Descending
	}
	val, read, err := decodeInt16FromSlice(b[1:], order)
	return val, read + 1, err
}

func DecodeInt16FromString(s string) (int16, int, error) {
	return DecodeInt16FromSlice(zbytes.StringToBytesUnsafe(s))
}

func decodeInt16FromSlice(b []byte, order Order) (int16, int, error) {
	if len(b) < 2 {
		return int16(0), 0, io.EOF
	}

	bslice := b[:2]
	if order == Descending {
		bslice = make([]byte, 2)
		copy(bslice, b[:2])
		xorbytes(bslice)
	}

	val := int16((bslice[0] ^ 0x80) & 0xff)
	val = (val << 8) + int16(bslice[1]&0xff)

	return val, 2, nil
}

// Encode a 32-bit integer
func EncodeInt32(val int32, order Order) []byte {
	var tmp = make([]byte, 5)
	tmp[0] = FixedInt32
	binary.BigEndian.PutUint32(tmp[1:], uint32(val)^(1<<31))
	if order == Descending {
		xorbytes(tmp)
	}
	return tmp
}

// Decode a 32-bit integer stored in the given buffer
func DecodeInt32(buffer io.ReadSeeker) (int32, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return int32(0), err
	}
	if count < 1 {
		return int32(0), io.EOF
	}
	code := tmp[0]
	if code != FixedInt32 && code != fixedInt32Desc {
		return 0, ErrWrongTypeCode
	}
	if code == fixedInt32Desc {
		return decodeInt32(buffer, Descending)
	}
	return decodeInt32(buffer, Ascending)
}

func decodeInt32(buffer io.ReadSeeker, order Order) (int32, error) {
	tmp := make([]byte, 4)
	count, err := buffer.Read(tmp)
	if err != nil {
		return int32(0), err
	}
	if count < 4 {
		return int32(0), io.EOF
	}
	val, _, err := decodeInt32FromSlice(tmp, order)

	return val, err
}

func DecodeInt32FromSlice(b []byte) (int32, int, error) {
	if len(b) < 1 {
		return int32(0), 0, io.EOF
	}
	code := b[0]
	if code != FixedInt32 && code != fixedInt32Desc {
		return 0, 0, ErrWrongTypeCode
	}
	order := Ascending
	if code == fixedInt32Desc {
		order = Descending
	}
	val, read, err := decodeInt32FromSlice(b[1:], order)
	return val, read + 1, err
}

func DecodeInt32FromString(s string) (int32, int, error) {
	return DecodeInt32FromSlice(zbytes.StringToBytesUnsafe(s))
}

func decodeInt32FromSlice(b []byte, order Order) (int32, int, error) {
	if len(b) < 4 {
		return int32(0), 0, io.EOF
	}

	bSlice := b[:4]
	if order == Descending {
		bSlice = make([]byte, 4)
		copy(bSlice, b[:4])
		xorbytes(bSlice)
	}

	val := int32((bSlice[0] ^ 0x80) & 0xff)
	for i := 1; i < 4; i++ {
		val = (val << 8) + int32(bSlice[i]&0xff)
	}

	return val, 4, nil
}

// Encode a 64-bit integer
func EncodeInt64(val int64, order Order) []byte {
	var tmp = make([]byte, 9)
	tmp[0] = FixedInt64
	binary.BigEndian.PutUint64(tmp[1:], uint64(val)^(1<<63))
	if order == Descending {
		xorbytes(tmp)
	}

	return tmp
}

// Decode a 64-bit integer stored in the given buffer
func DecodeInt64(buffer io.ReadSeeker) (int64, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return int64(0), err
	}
	if count < 1 {
		return int64(0), io.EOF
	}
	code := tmp[0]
	if code != FixedInt64 && code != fixedInt64Desc {
		return int64(0), ErrWrongTypeCode
	}
	if code == fixedInt64Desc {
		return decodeInt64(buffer, Descending)
	}
	return decodeInt64(buffer, Ascending)
}

func decodeInt64(buffer io.ReadSeeker, order Order) (int64, error) {
	tmp := make([]byte, 8)
	count, err := buffer.Read(tmp)
	if err != nil {
		return int64(0), err
	}
	if count < 8 {
		return int64(0), io.EOF
	}

	val, _, err := decodeInt64FromSlice(tmp, order)

	return val, err
}

func DecodeInt64FromSlice(b []byte) (int64, int, error) {
	if len(b) < 1 {
		return int64(0), 0, io.EOF
	}
	code := b[0]
	if code != FixedInt64 && code != fixedInt64Desc {
		return int64(0), 0, ErrWrongTypeCode
	}
	order := Ascending
	if code == fixedInt64Desc {
		order = Descending
	}
	val, read, err := decodeInt64FromSlice(b[1:], order)
	return val, read + 1, err
}

func DecodeInt64FromString(s string) (int64, int, error) {
	return DecodeInt64FromSlice(zbytes.StringToBytesUnsafe(s))
}

func decodeInt64FromSlice(b []byte, order Order) (int64, int, error) {
	if len(b) < 8 {
		return int64(0), 0, io.EOF
	}

	bSlice := b[:8]
	if order == Descending {
		bSlice = make([]byte, 8)
		copy(bSlice, b[:8])
		xorbytes(bSlice)
	}

	val := int64((bSlice[0] ^ 0x80) & 0xff)
	for i := 1; i < 8; i++ {
		val = (val << 8) + int64(bSlice[i]&0xff)
	}

	return val, 8, nil
}

///////////////////////////////////////////////////////////////////////////////
// Encode a floating-point type such that its encoded bytewise ordering matches
// its natural ordering.
//
///////////////////////////////////////////////////////////////////////////////

// Encode a 64 bit float
func EncodeFloat64(val float64, order Order) []byte {
	var tmp = make([]byte, 9)
	int64Val := int64(math.Float64bits(val))
	int64Val ^= (int64Val >> 63) | (-1 << 63)

	tmp[0] = FixedFloat64
	binary.BigEndian.PutUint64(tmp[1:], uint64(int64Val))

	if order == Descending {
		xorbytes(tmp)
	}

	return tmp
}

// Decode a 64-bit float stored in the given buffer
func DecodeFloat64(buffer io.ReadSeeker) (float64, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return float64(0), err
	}
	if count < 1 {
		return float64(0), io.EOF
	}
	code := tmp[0]
	if code != FixedFloat64 && code != fixedFloat64Desc {
		return float64(0), ErrWrongTypeCode
	}
	if code == fixedFloat64Desc {
		return decodeFloat64(buffer, Descending)
	}
	return decodeFloat64(buffer, Ascending)
}

func decodeFloat64(buffer io.ReadSeeker, order Order) (float64, error) {
	tmp := make([]byte, 8)
	count, err := buffer.Read(tmp)
	if err != nil {
		return float64(0), err
	}
	if count < 8 {
		return float64(0), io.EOF
	}
	val, _, err := decodeFloat64FromSlice(tmp, order)
	return val, err
}

func DecodeFloat64FromSlice(b []byte) (float64, int, error) {
	if len(b) < 1 {
		return float64(0), 0, io.EOF
	}
	code := b[0]
	if code != FixedFloat64 && code != fixedFloat64Desc {
		return float64(0), 0, ErrWrongTypeCode
	}
	order := Ascending
	if code == fixedFloat64Desc {
		order = Descending
	}
	val, read, err := decodeFloat64FromSlice(b[1:], order)
	return val, read + 1, err
}

func DecodeFloat64FromString(s string) (float64, int, error) {
	return DecodeFloat64FromSlice(zbytes.StringToBytesUnsafe(s))
}

func decodeFloat64FromSlice(b []byte, order Order) (float64, int, error) {
	if len(b) < 8 {
		return float64(0), 0, io.EOF
	}
	bSlice := b[:8]
	if order == Descending {
		bSlice = make([]byte, 8)
		copy(bSlice, b[:8])
		xorbytes(bSlice)
	}
	int64Val := int64(binary.BigEndian.Uint64(bSlice))
	int64Val ^= (^int64Val >> 63) | (-1 << 63)
	return math.Float64frombits(uint64(int64Val)), 8, nil
}

// Encode a 32 bit float
func EncodeFloat32(val float32, order Order) []byte {
	var tmp = make([]byte, 5)
	int32Val := int32(math.Float32bits(val))
	int32Val ^= (int32Val >> 31) | (-1 << 31)

	tmp[0] = FixedFloat32
	binary.BigEndian.PutUint32(tmp[1:], uint32(int32Val))

	if order == Descending {
		xorbytes(tmp)
	}

	return tmp
}

// Decode a 32-bit integer stored in the given buffer
func DecodeFloat32(buffer io.ReadSeeker) (float32, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return float32(0), err
	}
	if count < 1 {
		return float32(0), io.EOF
	}
	code := tmp[0]
	if code != FixedFloat32 && code != fixedFloat32Desc {
		return float32(0), ErrWrongTypeCode
	}
	if code == fixedFloat32Desc {
		return decodeFloat32(buffer, Descending)
	}
	return decodeFloat32(buffer, Ascending)
}

func decodeFloat32(buffer io.ReadSeeker, order Order) (float32, error) {
	tmp := make([]byte, 4)
	count, err := buffer.Read(tmp)
	if err != nil {
		return float32(0), err
	}
	if count < 4 {
		return float32(0), io.EOF
	}
	val, _, err := decodeFloat32FromSlice(tmp, order)
	return val, err
}

func DecodeFloat32FromSlice(b []byte) (float32, int, error) {
	if len(b) < 1 {
		return float32(0), 0, io.EOF
	}
	code := b[0]
	if code != FixedFloat32 && code != fixedFloat32Desc {
		return float32(0), 0, ErrWrongTypeCode
	}
	order := Ascending
	if code == fixedFloat32Desc {
		order = Descending
	}
	val, read, err := decodeFloat32FromSlice(b[1:], order)
	return val, read + 1, err
}

func DecodeFloat32FromString(s string) (float32, int, error) {
	return DecodeFloat32FromSlice(zbytes.StringToBytesUnsafe(s))
}

func decodeFloat32FromSlice(b []byte, order Order) (float32, int, error) {
	if len(b) < 4 {
		return float32(0), 0, io.EOF
	}
	bSlice := b[:4]
	if order == Descending {
		bSlice = make([]byte, 4)
		copy(bSlice, b[:4])
		xorbytes(bSlice)
	}
	int32Val := int32(binary.BigEndian.Uint32(bSlice))
	int32Val ^= (^int32Val >> 31) | (-1 << 31)
	return math.Float32frombits(uint32(int32Val)), 4, nil
}

///////////////////////////////////////////////////////////////////////////////
// Encode a variable-length string such that its encoded bytewise ordering matches
// its natural ordering.
///////////////////////////////////////////////////////////////////////////////

func EncodeString(val string, order Order) ([]byte, error) {
	if strings.Contains(val, string(StringTerminator)) {
		return nil, ErrStringContainsTerminator
	}

	buff := bytes.NewBuffer(make([]byte, 0, len(val)+2))
	buff.WriteByte(Text)
	buff.WriteString(val)
	buff.WriteByte(StringTerminator)
	result := buff.Bytes()
	if order == Descending {
		xorbytes(result)
	}
	return result, nil
}

// Decode a 32-bit integer stored in the given buffer
func DecodeString(buffer io.ReadSeeker) (string, error) {
	tmp := make([]byte, 1)
	count, err := buffer.Read(tmp)
	if err != nil {
		return "", err
	}
	if count < 1 {
		return "", io.EOF
	}
	code := tmp[0]
	if code != Text && code != textDesc {
		return "", ErrWrongTypeCode
	}
	if code == textDesc {
		return decodeString(buffer, Descending)
	}
	return decodeString(buffer, Ascending)
}

func decodeString(buffer io.ReadSeeker, order Order) (string, error) {
	var terminator byte
	if order == Ascending {
		terminator = StringTerminator
	} else {
		terminator = stringTerminatorDesc
	}

	strBytes := make([]byte, 0, 10)

	// Read 10-byte chunks until we find the terminator
	done := false
	i := 0
	count := 0
	var err error
	for !done {
		tmp := make([]byte, 10)
		count, err = buffer.Read(tmp)
		if err != nil && err != io.EOF {
			return "", err
		}
		for i = 0; i < count; i++ {
			if tmp[i] != terminator {
				strBytes = append(strBytes, tmp[i])
			} else {
				done = true
				break
			}
		}

		if err == io.EOF && !done {
			return "", io.EOF
		}
	}

	// We probably read past the terminator, reset the reader to the end of the string
	if i+1 < count {
		_, err := buffer.Seek(int64(i-count+1), io.SeekCurrent)
		if err != nil {
			return "", errors.Wrap(err, "error resetting seeker")
		}
	}

	if order == Descending {
		xorbytes(strBytes)
	}

	return string(strBytes), nil
}

func DecodeStringFromSlice(b []byte) (string, int, error) {
	if len(b) < 1 {
		return "", 0, io.EOF
	}
	code := b[0]
	if code != Text && code != textDesc {
		return "", 0, ErrWrongTypeCode
	}
	order := Ascending
	if code == textDesc {
		order = Descending
	}
	val, read, err := decodeStringFromSlice(b[1:], order)
	return val, read + 1, err
}

func DecodeStringFromString(s string) (string, int, error) {
	return DecodeStringFromSlice(zbytes.StringToBytesUnsafe(s))
}

func decodeStringFromSlice(b []byte, order Order) (string, int, error) {
	var terminator byte
	if order == Ascending {
		terminator = StringTerminator
	} else {
		terminator = stringTerminatorDesc
	}

	// Find the terminator
	length := -1
	for i, bite := range b {
		if bite == terminator {
			length = i
			break
		}
	}

	if length < 0 {
		return "", 0, io.EOF
	}

	strBytes := b[:length]

	if order == Descending {
		strBytes = make([]byte, length)
		copy(strBytes, b[:length])
		xorbytes(strBytes)
	}

	return string(strBytes), length + 1, nil
}

// Size returns the size of the next encoded value in the buffer. If there is
// no next encoded value in the buffer, it returns io.EOF.
func Size(buf []byte) (int, error) {
	if len(buf) == 0 {
		return 0, io.EOF
	}
	switch buf[0] {
	case Null, nullDesc:
		return 1, nil
	case FixedInt8, fixedInt8Desc:
		return 2, nil
	case FixedInt16, fixedInt16Desc:
		return 3, nil
	case FixedInt32, fixedInt32Desc, FixedFloat32, fixedFloat32Desc:
		return 5, nil
	case FixedInt64, fixedInt64Desc, FixedFloat64, fixedFloat64Desc:
		return 9, nil
	case Text:
		for i, b := range buf[1:] {
			if b == StringTerminator {
				return i + 2, nil
			}
		}
		return 0, io.EOF
	case textDesc:
		for i, b := range buf[1:] {
			if b == stringTerminatorDesc {
				return i + 2, nil
			}
		}
		return 0, io.EOF
	}
	return 0, ErrUnsupportedType
}
