package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type DataRequest uint16

const (
	DrPoll         DataRequest = 0x0000
	DrUrlProduct   DataRequest = 0x0001
	DrUrlUserGuide DataRequest = 0x0002
	DrUrlSupport   DataRequest = 0x0003
	DrUrlPersUdr   DataRequest = 0x0004
	DrUrlPersGdtf  DataRequest = 0x0005
)

type ArtDataRequest struct {
	EstaMan uint16
	Oem     uint16
	Request DataRequest
	spare   [22]byte
}

func (adr *ArtDataRequest) GetOpCode() uint16 {
	return OpDataRequest
}

func (adr *ArtDataRequest) UnmarshalBinary(data []byte) error {

	if len(data) < 40 {
		return errors.New("ArtDataRequest packet must be at least 40 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpDataRequest {
		return errors.New("packet does not have the correct OpCode for an ArtDataRequest packet")
	}

	offset := 12
	adr.EstaMan = binary.BigEndian.Uint16(data[offset : offset+2])
	adr.Oem = binary.BigEndian.Uint16(data[offset+2 : offset+4])
	adr.Request = DataRequest(binary.BigEndian.Uint16(data[offset+4 : offset+6]))
	copy(adr.spare[:], data[offset+6:offset+28])
	return nil
}

func (adr *ArtDataRequest) MarshalBinary() ([]byte, error) {
	data := make([]byte, 12+28)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpDataRequest)
	data[10] = 0
	data[11] = 14
	binary.BigEndian.PutUint16(data[12:14], adr.EstaMan)
	binary.BigEndian.PutUint16(data[14:16], adr.Oem)
	binary.BigEndian.PutUint16(data[16:18], uint16(adr.Request))
	copy(data[18:40], adr.spare[:])
	return data, nil
}
