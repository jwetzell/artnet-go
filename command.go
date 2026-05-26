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
	Length    uint16
	Data      []uint8
}

func (ap *ArtCommand) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtCommand) GetProtVer() uint16 {
	return uint16(ap.ProtVerHi)<<8 + uint16(ap.ProtVerLo)
}

func (ap *ArtCommand) GetID() [8]uint8 {
	return ap.ID
}

func (ap *ArtCommand) SetData(data string) error {
	if len(data) > 512 {
		return errors.New("data length must be less than or equal to 512 bytes")
	}
	ap.Data = make([]uint8, len(data))
	copy(ap.Data, data)
	ap.Length = uint16(len(data))
	return nil
}

func (ap *ArtCommand) UnmarshalBinary(data []byte) error {

	if len(data) < 16 {
		return errors.New("ArtCommand packet must be at least 16 bytes long")
	}

	copy(ap.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ap.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ap.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ap.ProtVerHi = data[10]
	ap.ProtVerLo = data[11]

	offset := 12
	ap.EstaMan = binary.LittleEndian.Uint16(data[offset : offset+2])
	ap.Length = binary.LittleEndian.Uint16(data[offset+2 : offset+4])

	dataLength := int(ap.Length)

	if len(data[offset+4:]) < dataLength {
		return errors.New("[]byte length not long enough to contain data length specified in packet")
	}
	ap.Data = make([]uint8, dataLength)
	copy(ap.Data, data[offset+4:offset+4+dataLength])
	return nil
}

func (ap *ArtCommand) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+8)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.ProtVerHi
	data[11] = ap.ProtVerLo
	binary.LittleEndian.PutUint16(data[12:14], ap.EstaMan)
	binary.LittleEndian.PutUint16(data[14:16], ap.Length)
	data = append(data, ap.Data...)
	return data, nil
}
