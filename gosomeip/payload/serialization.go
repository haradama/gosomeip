package payload

import (
	"encoding/binary"
	"errors"
	"math"
)

const (
	// Data type constants
	BoolType    = 0
	Uint8Type   = 1
	Uint16Type  = 2
	Uint32Type  = 3
	Uint64Type  = 4
	Sint8Type   = 5
	Sint16Type  = 6
	Sint32Type  = 7
	Sint64Type  = 8
	Float32Type = 9
	Float64Type = 10

	// Fixed header size
	HeaderSize = 1 // 1 byte for type
)

var (
	ErrInvalidPayload = errors.New("invalid payload")
)

// Serializer handles serialization
type Serializer struct{}

// NewSerializer creates a new Serializer
func NewSerializer() *Serializer {
	return &Serializer{}
}

// Serialize serializes a Go value into bytes
func (s *Serializer) Serialize(value interface{}) ([]byte, error) {

	switch v := value.(type) {

	case bool:
		buf := make([]byte, HeaderSize+1)
		buf[0] = BoolType
		if v {
			buf[1] = 1
		} else {
			buf[1] = 0
		}
		return buf, nil

	case uint8:
		return s.serializeUint(Uint8Type, uint64(v))

	case uint16:
		return s.serializeUint(Uint16Type, uint64(v))

	case uint32:
		return s.serializeUint(Uint32Type, uint64(v))

	case uint64:
		return s.serializeUint(Uint64Type, v)

	case int8:
		return s.serializeSint(Sint8Type, int64(v))

	case int16:
		return s.serializeSint(Sint16Type, int64(v))

	case int32:
		return s.serializeSint(Sint32Type, int64(v))

	case int64:
		return s.serializeSint(Sint64Type, v)

	case float32:
		bits := math.Float32bits(v)
		buf := make([]byte, HeaderSize+4)
		buf[0] = Float32Type
		binary.LittleEndian.PutUint32(buf[HeaderSize:], bits)
		return buf, nil

	case float64:
		bits := math.Float64bits(v)
		buf := make([]byte, HeaderSize+8)
		buf[0] = Float64Type
		binary.LittleEndian.PutUint64(buf[HeaderSize:], bits)
		return buf, nil

	default:
		return nil, ErrInvalidPayload
	}
}

// serializeUint serializes an unsigned integer value
func (s *Serializer) serializeUint(t uint8, x uint64) ([]byte, error) {
	buf := make([]byte, HeaderSize+8)
	buf[0] = t
	binary.BigEndian.PutUint64(buf[HeaderSize:], x)
	return buf, nil
}

// serializeSint serializes a signed integer value
func (s *Serializer) serializeSint(t uint8, x int64) ([]byte, error) {
	buf := make([]byte, HeaderSize+8)
	buf[0] = t
	binary.BigEndian.PutUint64(buf[HeaderSize:], uint64(x))
	return buf, nil
}

// SerializeArray serializes an array payload
func (s *Serializer) SerializeArray(payload []interface{}) ([]byte, error) {
	var buf []byte
	for _, item := range payload {
		b, err := s.Serialize(item)
		if err != nil {
			return nil, err
		}
		buf = append(buf, b...)
	}
	return buf, nil
}
