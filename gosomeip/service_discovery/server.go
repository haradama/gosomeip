// server.go

package service_discovery

import (
	"net"

	"github.com/my/somip/header"
)

// Server implements a SOME/IP SD server
type Server struct {
	// ...other fields
}

// OfferService sends an OfferService message to advertise a service instance.
func (s *Server) OfferService(serviceID uint16, instanceID uint16,
	majorVersion uint8, minorVersion uint32, ttl uint32) error {

	// Create SOME/IP SD header
	header := &header.Header{
		MessageType: header.MessageTypeNotification,
		// Set other header fields
	}

	// Create OfferService entry
	offerEntry := &sdEntry{
		Type:         sdEntryTypeOfferService,
		ServiceID:    serviceID,
		InstanceID:   instanceID,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		TTL:          ttl,
		// Set other entry fields
	}

	// Add IPv4 Endpoint option
	ipAddr := net.ParseIP("192.168.0.5")
	endpointOpt := &ipv4EndpointOption{
		IPAddress: ipAddr,
		L4Proto:   udpProto,
		Port:      5000,
	}

	// Create payload
	payload := []interface{}{offerEntry, endpointOpt}

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
