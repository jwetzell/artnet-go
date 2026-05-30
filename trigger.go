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
	id        [8]uint8
	opCode    uint16
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
	return OpTrigger
}

func (at *ArtTrigger) UnmarshalBinary(data []byte) error {

	if len(data) < 18 {
		return errors.New("ArtTrigger packet must be at least 18 bytes long")
	}

	copy(at.id[:], data[0:8])

	if !slices.Equal(ArtNetID[:], at.id[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	at.opCode = binary.LittleEndian.Uint16(data[8:10])
	if at.opCode != OpTrigger {
		return errors.New("packet does not have the correct OpCode for an ArtTrigger packet")
	}
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
	copy(data[0:8], at.id[:])
	binary.LittleEndian.PutUint16(data[8:10], at.opCode)
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
