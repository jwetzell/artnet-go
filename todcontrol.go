package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type TodControlCommand uint8

const (
	AtcNone   TodControlCommand = 0x00
	AtcFlush  TodControlCommand = 0x01
	AtcEnd    TodControlCommand = 0x02
	AtcIncOn  TodControlCommand = 0x03
	AtcIncOff TodControlCommand = 0x04
)

type ArtTodControl struct {
	filler1 uint8
	filler2 uint8
	spare1  uint8
	spare2  uint8
	spare3  uint8
	spare4  uint8
	spare5  uint8
	spare6  uint8
	spare7  uint8
	Net     uint8
	Command TodControlCommand
	Address uint8
}

func (adr *ArtTodControl) GetOpCode() uint16 {
	return OpTodControl
}

func (adr *ArtTodControl) UnmarshalBinary(data []byte) error {

	if len(data) < 24 {
		return errors.New("ArtTodControl packet must be at least 24 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpTodControl {
		return errors.New("packet does not have the correct OpCode for an ArtTodControl packet")
	}

	offset := 12
	adr.filler1 = data[offset]
	adr.filler2 = data[offset+1]
	adr.spare1 = data[offset+2]
	adr.spare2 = data[offset+3]
	adr.spare3 = data[offset+4]
	adr.spare4 = data[offset+5]
	adr.spare5 = data[offset+6]
	adr.spare6 = data[offset+7]
	adr.spare7 = data[offset+8]
	adr.Net = data[offset+9]
	adr.Command = TodControlCommand(data[offset+10])
	adr.Address = data[offset+11]
	return nil
}

func (adr *ArtTodControl) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+16)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpTodControl)
	data[10] = 0
	data[11] = 14
	data[12] = adr.filler1
	data[13] = adr.filler2
	data[14] = adr.spare1
	data[15] = adr.spare2
	data[16] = adr.spare3
	data[17] = adr.spare4
	data[18] = adr.spare5
	data[19] = adr.spare6
	data[20] = adr.spare7
	data[21] = adr.Net
	data[22] = uint8(adr.Command)
	data[23] = adr.Address
	return data, nil
}
