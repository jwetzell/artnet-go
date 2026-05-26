package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtSync struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	Aux1      uint8
	Aux2      uint8
}

func (as *ArtSync) GetOpCode() uint16 {
	return as.OpCode
}

func (as *ArtSync) GetProtVer() uint16 {
	return uint16(as.ProtVerHi)<<8 + uint16(as.ProtVerLo)
}

func (as *ArtSync) GetID() [8]uint8 {
	return as.ID
}

func (as *ArtSync) UnmarshalBinary(data []byte) error {

	if len(data) < 14 {
		return errors.New("ArtSync packet must be at least 14 bytes long")
	}

	copy(as.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], as.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	as.OpCode = binary.LittleEndian.Uint16(data[8:10])
	as.ProtVerHi = data[10]
	as.ProtVerLo = data[11]

	offset := 12
	as.Aux1 = data[offset]
	as.Aux2 = data[offset+1]
	return nil
}

func (as *ArtSync) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+6)
	copy(data[0:8], as.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], as.OpCode)
	data[10] = as.ProtVerHi
	data[11] = as.ProtVerLo
	data[12] = as.Aux1
	data[13] = as.Aux2
	return data, nil
}
