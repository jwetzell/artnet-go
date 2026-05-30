package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type Command string

const (
	SwoutText Command = "SwoutText"
	SwinText  Command = "SwinText"
)

type ArtCommand struct {
	EstaMan uint16
	Data    []uint8
}

func (ac *ArtCommand) GetOpCode() uint16 {
	return OpCommand
}

func (ac *ArtCommand) UnmarshalBinary(data []byte) error {

	if len(data) < 16 {
		return errors.New("ArtCommand packet must be at least 16 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpCommand {
		return errors.New("packet does not have the correct OpCode for an ArtCommand packet")
	}

	offset := 12
	ac.EstaMan = binary.BigEndian.Uint16(data[offset : offset+2])

	dataLength := int(binary.BigEndian.Uint16(data[offset+2 : offset+4]))

	if len(data[offset+4:]) < dataLength {
		return errors.New("[]byte length not long enough to contain data length specified in packet")
	}
	ac.Data = make([]uint8, dataLength)
	copy(ac.Data, data[offset+4:offset+4+dataLength])
	return nil
}

func (ac *ArtCommand) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+8)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpCommand)
	data[10] = 0
	data[11] = 14
	binary.BigEndian.PutUint16(data[12:14], ac.EstaMan)
	dataLength := uint16(len(ac.Data))
	if dataLength > 512 {
		return nil, errors.New("data length must be less than or equal to 512 bytes")
	}
	binary.BigEndian.PutUint16(data[14:16], dataLength)
	data = append(data, ac.Data...)
	return data, nil
}
