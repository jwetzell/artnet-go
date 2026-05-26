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

func (ap *ArtSync) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtSync) GetProtVer() uint16 {
	return uint16(ap.ProtVerHi)<<8 + uint16(ap.ProtVerLo)
}

func (ap *ArtSync) GetID() [8]uint8 {
	return ap.ID
}

func (ap *ArtSync) UnmarshalBinary(data []byte) error {

	if len(data) < 14 {
		return errors.New("ArtSync packet must be at least 14 bytes long")
	}

	copy(ap.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ap.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ap.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ap.ProtVerHi = data[10]
	ap.ProtVerLo = data[11]

	offset := 12
	ap.Aux1 = data[offset]
	ap.Aux2 = data[offset+1]
	return nil
}

func (ap *ArtSync) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+6)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.ProtVerHi
	data[11] = ap.ProtVerLo
	data[12] = ap.Aux1
	data[13] = ap.Aux2
	return data, nil
}
