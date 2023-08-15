package payload

import (
	"encoding/binary"
)

// Deserializer handles SOME/IP payload deserialization
type Deserializer struct{}

// NewDeserializer creates a new Deserializer
func NewDeserializer() *Deserializer {
	return &Deserializer{}
}

// Deserialize deserializes a payload from bytes into an interface{}
func (d *Deserializer) Deserialize(payload []byte) (interface{}, error) {
	if len(payload) < 4 {
		return nil, ErrInvalidPayload
	}

	var deserialized interface{}

	// Deserialize based on length
	switch len(payload) {
	case 4:
		i := int32(binary.BigEndian.Uint32(payload))
		deserialized = i
	case 8:
		f := binary.BigEndian.Uint64(payload)
		deserialized = float64(f)
	default:
		deserialized = string(payload)
	}

	return deserialized, nil
}

// DeserializeArray deserializes a payload into an array
func (d *Deserializer) DeserializeArray(payload []byte) ([]interface{}, error) {
	var deserialized []interface{}

	for len(payload) > 0 {
		item, err := d.Deserialize(payload)
		if err != nil {
			return nil, err
		}

		deserialized = append(deserialized, item)

		// Advance payload
		if str, ok := item.(string); ok {
			payload = payload[len(str):]
		} else {
			payload = payload[4:]
		}
	}

	return deserialized, nil
}
