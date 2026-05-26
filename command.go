package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type Command string

const (
	SwoutText Command = "SwoutText"
	SwinText  Command = "SwinText"
)

type ArtCommand struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	EstaMan   uint16
	Data      []uint8
}

func (ac *ArtCommand) GetOpCode() uint16 {
	return ac.OpCode
}

func (ac *ArtCommand) GetProtVer() uint16 {
	return uint16(ac.ProtVerHi)<<8 + uint16(ac.ProtVerLo)
}

func (ac *ArtCommand) GetID() [8]uint8 {
	return ac.ID
}

func (ac *ArtCommand) Length() uint16 {
	return uint16(len(ac.Data))
}

func (ac *ArtCommand) UnmarshalBinary(data []byte) error {

	if len(data) < 16 {
		return errors.New("ArtCommand packet must be at least 16 bytes long")
	}

	copy(ac.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ac.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ac.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ac.ProtVerHi = data[10]
	ac.ProtVerLo = data[11]

	offset := 12
	ac.EstaMan = binary.LittleEndian.Uint16(data[offset : offset+2])

	dataLength := int(binary.LittleEndian.Uint16(data[offset+2 : offset+4]))

	if len(data[offset+4:]) < dataLength {
		return errors.New("[]byte length not long enough to contain data length specified in packet")
	}
	ac.Data = make([]uint8, dataLength)
	copy(ac.Data, data[offset+4:offset+4+dataLength])
	return nil
}

func (ac *ArtCommand) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+8)
	copy(data[0:8], ac.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ac.OpCode)
	data[10] = ac.ProtVerHi
	data[11] = ac.ProtVerLo
	binary.LittleEndian.PutUint16(data[12:14], ac.EstaMan)
	dataLength := uint16(len(ac.Data))
	if dataLength > 512 {
		return nil, errors.New("data length must be less than or equal to 512 bytes")
	}
	binary.LittleEndian.PutUint16(data[14:16], dataLength)
	data = append(data, ac.Data...)
	return data, nil
}
