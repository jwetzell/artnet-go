package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type Key uint8

const (
	KeyASCII Key = 0
	KeyMacro Key = 1
	KeySoft  Key = 2
	KeyShow  Key = 3
)

type ArtTrigger struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	filler1   uint8
	filler2   uint8
	Oem       uint16
	Key       Key
	SubKey    uint8
	Data      []uint8
}

func (at *ArtTrigger) GetOpCode() uint16 {
	return at.OpCode
}

func (at *ArtTrigger) GetProtVer() uint16 {
	return uint16(at.ProtVerHi)<<8 + uint16(at.ProtVerLo)
}

func (at *ArtTrigger) GetID() [8]uint8 {
	return at.ID
}

func (at *ArtTrigger) UnmarshalBinary(data []byte) error {

	if len(data) < 18 {
		return errors.New("ArtTrigger packet must be at least 18 bytes long")
	}

	copy(at.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], at.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	at.OpCode = binary.LittleEndian.Uint16(data[8:10])
	at.ProtVerHi = data[10]
	at.ProtVerLo = data[11]

	offset := 12
	at.filler1 = data[offset]
	at.filler2 = data[offset+1]
	at.Oem = binary.LittleEndian.Uint16(data[offset+2 : offset+4])
	at.Key = Key(data[offset+4])
	at.SubKey = data[offset+5]
	at.Data = make([]uint8, len(data[offset+6:]))
	copy(at.Data, data[offset+6:])
	return nil
}

func (at *ArtTrigger) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+10)
	copy(data[0:8], at.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], at.OpCode)
	data[10] = at.ProtVerHi
	data[11] = at.ProtVerLo
	data[12] = at.filler1
	data[13] = at.filler2
	binary.LittleEndian.PutUint16(data[14:16], at.Oem)
	data[16] = uint8(at.Key)
	data[17] = at.SubKey
	data = append(data, at.Data...)
	return data, nil
}
