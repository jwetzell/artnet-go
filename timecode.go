package artnet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"slices"
)

type ArtTimeCode struct {
	ID        []uint8
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

func NewArtTimeCode(data []byte) (*ArtTimeCode, error) {
	artTimeCode := ArtTimeCode{}

	err := artTimeCode.UnmarshalBinary(data)

	if err != nil {
		return nil, err
	}

	return &artTimeCode, nil
}

func (atc *ArtTimeCode) GetOpCode() uint16 {
	return atc.OpCode
}

func (atc *ArtTimeCode) GetProtVer() uint16 {
	return uint16(atc.ProtVerHi)<<8 + uint16(atc.ProtVerLo)
}

func (atc *ArtTimeCode) GetID() []uint8 {
	return atc.ID
}

func (atc *ArtTimeCode) UnmarshalBinary(data []byte) error {
	fmt.Println(data)
	if len(data) < 14 {
		return errors.New("ArtTimeCode packet must be at least 14 bytes long")
	}

	atc.ID = data[0:8]

	if !slices.Equal(ArtNetID, atc.ID) {
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
	data := []byte(ArtNetID)
	data = append(data, byte(atc.OpCode), byte(atc.OpCode>>8), atc.ProtVerHi, atc.ProtVerLo, atc.Filler1, atc.StreamId, atc.Frames, atc.Seconds, atc.Minutes, atc.Hours, atc.Type)
	return data, nil
}
