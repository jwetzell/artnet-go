package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type DataRequest uint16

const (
	DrPoll         = 0x0000
	DrUrlProduct   = 0x0001
	DrUrlUserGuide = 0x0002
	DrUrlSupport   = 0x0003
	DrUrlPersUdr   = 0x0004
	DrUrlPersGdtf  = 0x0005
)

type ArtDataRequest struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	EstaMan   uint16
	Oem       uint16
	Request   uint16
	spare     [22]byte
}

func (ap *ArtDataRequest) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtDataRequest) GetProtVer() uint16 {
	return uint16(ap.ProtVerHi)<<8 + uint16(ap.ProtVerLo)
}

func (ap *ArtDataRequest) GetID() [8]uint8 {
	return ap.ID
}

func (ap *ArtDataRequest) UnmarshalBinary(data []byte) error {

	if len(data) < 40 {
		return errors.New("ArtDataRequest packet must be at least 40 bytes long")
	}

	copy(ap.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ap.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ap.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ap.ProtVerHi = data[10]
	ap.ProtVerLo = data[11]

	offset := 12
	ap.EstaMan = binary.LittleEndian.Uint16(data[offset : offset+2])
	ap.Oem = binary.LittleEndian.Uint16(data[offset+2 : offset+4])
	ap.Request = binary.LittleEndian.Uint16(data[offset+4 : offset+6])
	copy(ap.spare[:], data[offset+6:offset+28])
	return nil
}

func (ap *ArtDataRequest) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+32)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.ProtVerHi
	data[11] = ap.ProtVerLo
	binary.LittleEndian.PutUint16(data[12:14], ap.EstaMan)
	binary.LittleEndian.PutUint16(data[14:16], ap.Oem)
	binary.LittleEndian.PutUint16(data[16:18], ap.Request)
	copy(data[18:40], ap.spare[:])
	return data, nil
}
