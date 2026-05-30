package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type TodRequestCommand uint8

const (
	TodFull TodRequestCommand = 0x00
)

type ArtTodRequest struct {
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
	Command TodRequestCommand
	Address []uint8
}

func (adr *ArtTodRequest) GetOpCode() uint16 {
	return OpTodRequest
}

func (adr *ArtTodRequest) UnmarshalBinary(data []byte) error {

	if len(data) < 32 {
		return errors.New("ArtTodRequest packet must be at least 32 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpTodRequest {
		return errors.New("packet does not have the correct OpCode for an ArtTodRequest packet")
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
	adr.Command = TodRequestCommand(data[offset+10])
	addCount := data[offset+11]
	adr.Address = make([]uint8, addCount)
	if len(data[offset+12:]) < int(addCount) {
		return errors.New("[]byte length not long enough to contain address length specified in packet")
	}
	copy(adr.Address, data[offset+12:offset+12+int(addCount)])
	return nil
}

func (adr *ArtTodRequest) MarshalBinary() ([]byte, error) {
	if len(adr.Address) > 32 {
		return nil, errors.New("address count must not be greater than 32")
	}
	data := make([]byte, 8+16+len(adr.Address))
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpTodRequest)
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
	data[23] = uint8(len(adr.Address))
	copy(data[24:], adr.Address)
	return data, nil
}
