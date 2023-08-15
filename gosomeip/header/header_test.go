package header

import (
	"reflect"
	"testing"
)

func TestHeader_Serialize(t *testing.T) {
	type fields struct {
		MessageID        uint32
		Length           uint32
		RequestID        uint32
		ProtocolVersion  uint8
		InterfaceVersion uint8
		MessageType      uint8
		ReturnCode       uint8
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				MessageID:        tt.fields.MessageID,
				Length:           tt.fields.Length,
				RequestID:        tt.fields.RequestID,
				ProtocolVersion:  tt.fields.ProtocolVersion,
				InterfaceVersion: tt.fields.InterfaceVersion,
				MessageType:      tt.fields.MessageType,
				ReturnCode:       tt.fields.ReturnCode,
			}
			got, err := h.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Header.Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeserializeHeader(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Header
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeserializeHeader(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeserializeHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeserializeHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
