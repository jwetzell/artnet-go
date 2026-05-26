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
	EstaManHi uint8
	EstaManLo uint8
	OemHi     uint8
	OemLo     uint8
	RequestHi uint8
	RequestLo uint8
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

func (ap *ArtDataRequest) GetDataRequestType() DataRequest {
	return DataRequest(uint16(ap.RequestHi)<<8 + uint16(ap.RequestLo))
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
	ap.EstaManHi = data[offset]
	ap.EstaManLo = data[offset+1]
	ap.OemHi = data[offset+2]
	ap.OemLo = data[offset+3]
	ap.RequestHi = data[offset+4]
	ap.RequestLo = data[offset+5]
	copy(ap.spare[:], data[offset+6:offset+28])
	return nil
}

func (ap *ArtDataRequest) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+32)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.ProtVerHi
	data[11] = ap.ProtVerLo
	data[12] = ap.EstaManHi
	data[13] = ap.EstaManLo
	data[14] = ap.OemHi
	data[15] = ap.OemLo
	data[16] = ap.RequestHi
	data[17] = ap.RequestLo
	copy(data[18:40], ap.spare[:])
	return data, nil
}
