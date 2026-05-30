package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type DiagPriority uint8

const (
	DpLow      DiagPriority = 0x10
	DpMed      DiagPriority = 0x40
	DpHigh     DiagPriority = 0x80
	DpCritical DiagPriority = 0xe0
	DpVolatile DiagPriority = 0xf0
)

type ArtDiagData struct {
	filler1      uint8
	DiagPriority DiagPriority
	LogicalPort  uint8
	filler3      uint8
	Data         []uint8
}

func (add *ArtDiagData) GetOpCode() uint16 {
	return OpDiagData
}

func (add *ArtDiagData) UnmarshalBinary(data []byte) error {

	if len(data) < 18 {
		return errors.New("ArtDiagData packet must be at least 18 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpDiagData {
		return errors.New("packet does not have the correct OpCode for an ArtDiagData packet")
	}

	offset := 12
	add.filler1 = data[offset]
	add.DiagPriority = DiagPriority(data[offset+1])
	add.LogicalPort = data[offset+2]
	add.filler3 = data[offset+3]
	dataLength := binary.BigEndian.Uint16(data[offset+4 : offset+6])

	if len(data[offset+6:]) < int(dataLength) {
		return errors.New("[]byte length not long enough to contain data length specified in packet")
	}
	add.Data = make([]uint8, dataLength)
	copy(add.Data, data[offset+6:offset+6+int(dataLength)])
	return nil
}

func (add *ArtDiagData) MarshalBinary() ([]byte, error) {
	data := make([]byte, 12+6+len(add.Data))
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpDiagData)
	data[10] = 0
	data[11] = 14
	data[12] = add.filler1
	data[13] = uint8(add.DiagPriority)
	data[14] = add.LogicalPort
	data[15] = add.filler3
	dataLength := uint16(len(add.Data))
	if dataLength > 512 {
		return nil, errors.New("data length must be less than or equal to 512 bytes")
	}
	binary.BigEndian.PutUint16(data[16:18], dataLength)
	copy(data[18:], add.Data)
	return data, nil
}
