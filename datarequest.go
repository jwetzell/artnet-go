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

func (adr *ArtDataRequest) GetOpCode() uint16 {
	return adr.OpCode
}

func (adr *ArtDataRequest) GetProtVer() uint16 {
	return uint16(adr.ProtVerHi)<<8 + uint16(adr.ProtVerLo)
}

func (adr *ArtDataRequest) GetID() [8]uint8 {
	return adr.ID
}

func (adr *ArtDataRequest) UnmarshalBinary(data []byte) error {

	if len(data) < 40 {
		return errors.New("ArtDataRequest packet must be at least 40 bytes long")
	}

	copy(adr.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], adr.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	adr.OpCode = binary.LittleEndian.Uint16(data[8:10])
	adr.ProtVerHi = data[10]
	adr.ProtVerLo = data[11]

	offset := 12
	adr.EstaMan = binary.BigEndian.Uint16(data[offset : offset+2])
	adr.Oem = binary.BigEndian.Uint16(data[offset+2 : offset+4])
	adr.Request = binary.BigEndian.Uint16(data[offset+4 : offset+6])
	copy(adr.spare[:], data[offset+6:offset+28])
	return nil
}

func (adr *ArtDataRequest) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+32)
	copy(data[0:8], adr.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], adr.OpCode)
	data[10] = adr.ProtVerHi
	data[11] = adr.ProtVerLo
	binary.BigEndian.PutUint16(data[12:14], adr.EstaMan)
	binary.BigEndian.PutUint16(data[14:16], adr.Oem)
	binary.BigEndian.PutUint16(data[16:18], adr.Request)
	copy(data[18:40], adr.spare[:])
	return data, nil
}
