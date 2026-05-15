package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtTimeCode struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	Filler1   uint8
	StreamId  uint8
	Frames    uint8
	Seconds   uint8
	Minutes   uint8
	Hours     uint8
	Type      uint8
}

func (atc *ArtTimeCode) GetOpCode() uint16 {
	return atc.OpCode
}

func (atc *ArtTimeCode) GetProtVer() uint16 {
	return uint16(atc.ProtVerHi)<<8 + uint16(atc.ProtVerLo)
}

func (atc *ArtTimeCode) GetID() [8]uint8 {
	return atc.ID
}

func (atc *ArtTimeCode) UnmarshalBinary(data []byte) error {
	if len(data) < 14 {
		return errors.New("ArtTimeCode packet must be at least 14 bytes long")
	}

	copy(atc.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], atc.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	atc.OpCode = binary.LittleEndian.Uint16(data[8:10])
	atc.ProtVerHi = data[10]
	atc.ProtVerLo = data[11]

	offset := 12
	atc.Filler1 = data[offset]
	atc.StreamId = data[offset+1]
	atc.Frames = data[offset+2]
	atc.Seconds = data[offset+3]
	atc.Minutes = data[offset+4]
	atc.Hours = data[offset+5]
	atc.Type = data[offset+6]
	return nil
}

func (atc *ArtTimeCode) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+11)
	copy(data[0:8], atc.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], atc.OpCode)
	data[10] = atc.ProtVerHi
	data[11] = atc.ProtVerLo
	offset := 12
	data[offset] = atc.Filler1
	data[offset+1] = atc.StreamId
	data[offset+2] = atc.Frames
	data[offset+3] = atc.Seconds
	data[offset+4] = atc.Minutes
	data[offset+5] = atc.Hours
	data[offset+6] = atc.Type
	return data, nil
}
