package header

import (
	"encoding/binary"
	"errors"
)

// +-------------------------------------------------------------------------------+
// |                                   Message ID                                  |
// |                                    [32 Bit]                                   |
// +-------------------------------------------------------------------------------+
// |                                     Length                                    |
// |                                    [32 Bit]                                   |
// +-------------------------------------------------------------------------------+
// |                                   Request ID                                  |
// |                                    [32 Bit]                                   |
// +---------------------------------------+---------------------------------------+
// | Protocol Version | Interface Version  |   Message Type    |    Return Code    |
// |     [8 Bit]      |      [8 Bit]       |     [8 Bit]       |      [8 Bit]      |
// +---------------------------------------+---------------------------------------+

const (
	// ProtocolVersion is the SOME/IP protocol version
	ProtocolVersion uint8 = 0x01

	// InterfaceVersion is the SOME/IP interface version
	InterfaceVersion uint8 = 0x01

	// MessageTypeRequest is the request message type
	MessageTypeRequest uint8 = 0x00

	// MessageTypeIDResponse is the response message type
	MessageTypeResponse uint8 = 0x80

	// MessageTypeError is the error message type
	MessageTypeError uint8 = 0x81
)

// Error types
var (
	ErrInvalidHeader = errors.New("invalid header")
)

// Header represents the SOME/IP header
type Header struct {
	MessageID        uint32
	Length           uint32
	RequestID        uint32
	ProtocolVersion  uint8
	InterfaceVersion uint8
	MessageType      uint8
	ReturnCode       uint8
}

// Serialize serializes the header into bytes
func (h *Header) Serialize() ([]byte, error) {
	buf := make([]byte, 16)

	binary.BigEndian.PutUint32(buf[0:4], h.MessageID)
	binary.BigEndian.PutUint32(buf[4:8], h.Length)
	binary.BigEndian.PutUint32(buf[8:12], h.RequestID)
	buf[12] = h.ProtocolVersion
	buf[13] = h.InterfaceVersion
	buf[14] = h.MessageType
	buf[15] = h.ReturnCode

	return buf, nil
}

// DeserializeHeader deserializes a byte array into a Header
func DeserializeHeader(buf []byte) (*Header, error) {
	if len(buf) < 16 {
		return nil, ErrInvalidHeader
	}

	h := &Header{}

	h.MessageID = binary.BigEndian.Uint32(buf[0:4])
	h.Length = binary.BigEndian.Uint32(buf[4:8])
	h.RequestID = binary.BigEndian.Uint32(buf[8:12])
	h.ProtocolVersion = buf[12]
	h.InterfaceVersion = buf[13]
	h.MessageType = buf[14]
	h.ReturnCode = buf[15]

	return h, nil
}
