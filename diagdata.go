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
	ID           [8]uint8
	OpCode       uint16
	ProtVerHi    uint8
	ProtVerLo    uint8
	filler1      uint8
	DiagPriority DiagPriority
	LogicalPort  uint8
	filler3      uint8
	Data         []uint8
}

func (add *ArtDiagData) GetOpCode() uint16 {
	return add.OpCode
}

func (add *ArtDiagData) GetProtVer() uint16 {
	return uint16(add.ProtVerHi)<<8 + uint16(add.ProtVerLo)
}

func (add *ArtDiagData) GetID() [8]uint8 {
	return add.ID
}

func (add *ArtDiagData) Length() uint16 {
	return uint16(len(add.Data))
}

func (add *ArtDiagData) UnmarshalBinary(data []byte) error {

	if len(data) < 18 {
		return errors.New("ArtDiagData packet must be at least 18 bytes long")
	}

	copy(add.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], add.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	add.OpCode = binary.LittleEndian.Uint16(data[8:10])
	add.ProtVerHi = data[10]
	add.ProtVerLo = data[11]

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
	data := make([]byte, 8+10)
	copy(data[0:8], add.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], add.OpCode)
	data[10] = add.ProtVerHi
	data[11] = add.ProtVerLo
	data[12] = add.filler1
	data[13] = uint8(add.DiagPriority)
	data[14] = add.LogicalPort
	data[15] = add.filler3
	dataLength := uint16(len(add.Data))
	if dataLength > 512 {
		return nil, errors.New("data length must be less than or equal to 512 bytes")
	}
	binary.BigEndian.PutUint16(data[16:18], dataLength)
	data = append(data, add.Data...)
	return data, nil
}
