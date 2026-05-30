package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtSync struct {
	Aux1 uint8
	Aux2 uint8
}

func (as *ArtSync) GetOpCode() uint16 {
	return OpSync
}

func (as *ArtSync) UnmarshalBinary(data []byte) error {

	if len(data) < 14 {
		return errors.New("ArtSync packet must be at least 14 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpSync {
		return errors.New("packet does not have the correct OpCode for an ArtSync packet")
	}

	offset := 12
	as.Aux1 = data[offset]
	as.Aux2 = data[offset+1]
	return nil
}

func (as *ArtSync) MarshalBinary() ([]byte, error) {
	data := make([]byte, 12+2)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpSync)
	data[10] = 0
	data[11] = 14
	data[12] = as.Aux1
	data[13] = as.Aux2
	return data, nil
}
