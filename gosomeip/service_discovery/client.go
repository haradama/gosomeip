// client.go

package service_discovery

import (
	"net"

	"gosomeip/header"
)

// Client implements a SOME/IP SD client
type Client struct {
	// ...other fields
}

// FindService sends a FindService message to discover service instances
func (c *Client) FindService(serviceID uint16, instanceID uint16) error {

	// Create SOME/IP SD header
	header := &header.Header{
		MessageType: header.MessageTypeRequest,
		// Set other fields
	}

	// Create FindService entry
	findEntry := &sdEntry{
		Type:       sdEntryTypeFindService,
		ServiceID:  serviceID,
		InstanceID: instanceID,
		// Set other fields
	}

	// Create SD message payload
	payload := []interface{}{findEntry}

	// Serialize payload
	serializer := payload.NewSerializer()
	payloadBytes, err := serializer.SerializeArray(payload)
	if err != nil {
		return err
	}

	// Send SOME/IP SD message via UDP
	addr := &net.UDPAddr{
		IP:   net.ParseIP("224.224.224.244"),
		Port: 30490,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	_, err = conn.Write(header.Serialize())
	if err != nil {
		return err
	}

	_, err = conn.Write(payloadBytes)
	if err != nil {
		return err
	}

	return nil
}
