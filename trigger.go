package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtTrigger struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	filler1   uint8
	filler2   uint8
	OemHi     uint8
	OemLo     uint8
	Key       uint8
	SubKey    uint8
	Data      []uint8
}

func (ap *ArtTrigger) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtTrigger) GetProtVer() uint16 {
	return uint16(ap.ProtVerHi)<<8 + uint16(ap.ProtVerLo)
}

func (ap *ArtTrigger) GetID() [8]uint8 {
	return ap.ID
}

func (ap *ArtTrigger) UnmarshalBinary(data []byte) error {

	if len(data) < 18 {
		return errors.New("ArtTrigger packet must be at least 18 bytes long")
	}

	copy(ap.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ap.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ap.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ap.ProtVerHi = data[10]
	ap.ProtVerLo = data[11]

	offset := 12
	ap.filler1 = data[offset]
	ap.filler2 = data[offset+1]
	ap.OemHi = data[offset+2]
	ap.OemLo = data[offset+3]
	ap.Key = data[offset+4]
	ap.SubKey = data[offset+5]
	ap.Data = make([]uint8, len(data[offset+6:]))
	copy(ap.Data, data[offset+6:])
	return nil
}

func (ap *ArtTrigger) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+10)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.ProtVerHi
	data[11] = ap.ProtVerLo
	data[12] = ap.filler1
	data[13] = ap.filler2
	data[14] = ap.OemHi
	data[15] = ap.OemLo
	data[16] = ap.Key
	data[17] = ap.SubKey
	data = append(data, ap.Data...)
	return data, nil
}
