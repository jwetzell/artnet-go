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
	filler1 uint8
	filler2 uint8
	Oem     uint16
	Key     Key
	SubKey  uint8
	Data    [512]uint8
}

func (at *ArtTrigger) GetOpCode() uint16 {
	return OpTrigger
}

func (at *ArtTrigger) UnmarshalBinary(data []byte) error {

	if len(data) < 530 {
		return errors.New("ArtTrigger packet must be at least 530 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpTrigger {
		return errors.New("packet does not have the correct OpCode for an ArtTrigger packet")
	}

	offset := 12
	at.filler1 = data[offset]
	at.filler2 = data[offset+1]
	at.Oem = binary.LittleEndian.Uint16(data[offset+2 : offset+4])
	at.Key = Key(data[offset+4])
	at.SubKey = data[offset+5]
	copy(at.Data[:], data[offset+6:])
	return nil
}

func (at *ArtTrigger) MarshalBinary() ([]byte, error) {
	data := make([]byte, 12+6+512)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpTrigger)
	data[10] = 0
	data[11] = 14
	data[12] = at.filler1
	data[13] = at.filler2
	binary.LittleEndian.PutUint16(data[14:16], at.Oem)
	data[16] = uint8(at.Key)
	data[17] = at.SubKey
	copy(data[18:], at.Data[:])
	return data, nil
}
