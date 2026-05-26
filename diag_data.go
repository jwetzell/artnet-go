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
	LengthHi     uint8
	LengthLo     uint8
	Data         []uint8
}

func (ap *ArtDiagData) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtDiagData) GetProtVer() uint16 {
	return uint16(ap.ProtVerHi)<<8 + uint16(ap.ProtVerLo)
}

func (ap *ArtDiagData) GetID() [8]uint8 {
	return ap.ID
}

func (ap *ArtDiagData) UnmarshalBinary(data []byte) error {

	if len(data) < 18 {
		return errors.New("ArtDiagData packet must be at least 18 bytes long")
	}

	copy(ap.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ap.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ap.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ap.ProtVerHi = data[10]
	ap.ProtVerLo = data[11]

	offset := 12
	ap.filler1 = data[offset]
	ap.DiagPriority = DiagPriority(data[offset+1])
	ap.LogicalPort = data[offset+2]
	ap.filler3 = data[offset+3]
	ap.LengthHi = data[offset+4]
	ap.LengthLo = data[offset+5]

	dataLength := int(ap.LengthHi)<<8 + int(ap.LengthLo)

	if len(data[offset+6:]) < dataLength {
		return errors.New("[]byte length not long enough to contain data length specified in packet")
	}
	ap.Data = make([]uint8, dataLength)
	copy(ap.Data, data[offset+6:offset+6+dataLength])
	return nil
}

func (ap *ArtDiagData) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+10)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.ProtVerHi
	data[11] = ap.ProtVerLo
	data[12] = ap.filler1
	data[13] = uint8(ap.DiagPriority)
	data[14] = ap.LogicalPort
	data[15] = ap.filler3
	data[16] = ap.LengthHi
	data[17] = ap.LengthLo
	data = append(data, ap.Data...)
	return data, nil
}
