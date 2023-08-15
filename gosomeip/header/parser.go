package header

import (
	"errors"
)

var (
	ErrInvalidPacket = errors.New("invalid SOME/IP packet")
)

// Parser handles parsing of SOME/IP packets
type Parser struct {
	header  *Header
	payload []byte
}

// NewParser creates a new parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses a SOME/IP packet
func (p *Parser) Parse(packet []byte) error {
	// Validate overall packet length
	if len(packet) < 16 {
		return ErrInvalidPacket
	}

	// Deserialize header
	header, err := DeserializeHeader(packet[:16])
	if err != nil {
		return err
	}
	p.header = header

	// Get payload bytes
	payloadLen := len(packet) - 16
	if int(p.header.Length) < payloadLen {
		return ErrInvalidPacket
	}

	p.payload = packet[16:p.header.Length]

	return nil
}

// Header returns the parsed header
func (p *Parser) Header() *Header {
	return p.header
}

// Payload returns the parsed payload bytes
func (p *Parser) Payload() []byte {
	return p.payload
}
