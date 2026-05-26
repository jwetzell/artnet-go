package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtCommand struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	EstaManHi uint8
	EstaManLo uint8
	LengthHi  uint8
	LengthLo  uint8
	Data      []uint8
}

func (ap *ArtCommand) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtCommand) GetProtVer() uint16 {
	return uint16(ap.ProtVerHi)<<8 + uint16(ap.ProtVerLo)
}

func (ap *ArtCommand) GetID() [8]uint8 {
	return ap.ID
}

func (ap *ArtCommand) SetData(data string) error {
	if len(data) > 512 {
		return errors.New("data length must be less than or equal to 512 bytes")
	}
	ap.Data = make([]uint8, len(data))
	copy(ap.Data, data)
	ap.LengthHi = uint8(len(data) >> 8)
	ap.LengthLo = uint8(len(data) & 0xff)
	return nil
}

func (ap *ArtCommand) UnmarshalBinary(data []byte) error {

	if len(data) < 16 {
		return errors.New("ArtCommand packet must be at least 16 bytes long")
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
	ap.LengthHi = data[offset+2]
	ap.LengthLo = data[offset+3]

	dataLength := int(ap.LengthHi)<<8 + int(ap.LengthLo)

	if len(data[offset+4:]) < dataLength {
		return errors.New("[]byte length not long enough to contain data length specified in packet")
	}
	ap.Data = make([]uint8, dataLength)
	copy(ap.Data, data[offset+4:offset+4+dataLength])
	return nil
}

func (ap *ArtCommand) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+8)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.ProtVerHi
	data[11] = ap.ProtVerLo
	data[12] = ap.EstaManHi
	data[13] = ap.EstaManLo
	data[14] = ap.LengthHi
	data[15] = ap.LengthLo
	data = append(data, ap.Data...)
	return data, nil
}
