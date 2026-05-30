package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtTimeCode struct {
	filler1  uint8
	StreamId uint8
	Frames   uint8
	Seconds  uint8
	Minutes  uint8
	Hours    uint8
	Type     uint8
}

func (atc *ArtTimeCode) GetOpCode() uint16 {
	return OpTimeCode
}

func (atc *ArtTimeCode) UnmarshalBinary(data []byte) error {
	if len(data) < 19 {
		return errors.New("ArtTimeCode packet must be at least 14 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpTimeCode {
		return errors.New("packet does not have the correct OpCode for an ArtTimeCode packet")
	}

	offset := 12
	atc.filler1 = data[offset]
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
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpTimeCode)
	data[10] = 0
	data[11] = 14
	offset := 12
	data[offset] = atc.filler1
	data[offset+1] = atc.StreamId
	data[offset+2] = atc.Frames
	data[offset+3] = atc.Seconds
	data[offset+4] = atc.Minutes
	data[offset+5] = atc.Hours
	data[offset+6] = atc.Type
	return data, nil
}
