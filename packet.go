package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtNetPacket interface {
	GetOpCode() uint16
}

var ArtNetID []uint8 = []uint8{'A', 'r', 't', '-', 'N', 'e', 't', 0x00}

type ArtNetHeader struct {
	ID        []uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
}

func NewHeader(data []byte) (*ArtNetHeader, error) {
	if len(data) < 12 {
		return nil, errors.New("header must be at least 12 bytes")
	}

	if !slices.Equal(ArtNetID, data[0:8]) {
		return nil, errors.New("header id does not match Art-Net ID")
	}

	return &ArtNetHeader{
		OpCode:    binary.LittleEndian.Uint16(data[8:10]),
		ProtVerHi: data[10],
		ProtVerLo: data[11],
	}, nil

}
