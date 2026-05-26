package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtNzs struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	Sequence  uint8
	StartCode uint8
	SubUni    uint8
	Net       uint8
	Data      []uint8
}

func (an *ArtNzs) GetOpCode() uint16 {
	return an.OpCode
}

func (an *ArtNzs) GetProtVer() uint16 {
	return uint16(an.ProtVerHi)<<8 + uint16(an.ProtVerLo)
}

func (an *ArtNzs) GetID() [8]uint8 {
	return an.ID
}

func (an *ArtNzs) Length() uint16 {
	return uint16(len(an.Data))
}

func (an *ArtNzs) UnmarshalBinary(data []byte) error {

	if len(data) < 18 {
		return errors.New("ArtNzs packet must be at least 18 bytes long")
	}

	copy(an.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], an.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	an.OpCode = binary.LittleEndian.Uint16(data[8:10])
	an.ProtVerHi = data[10]
	an.ProtVerLo = data[11]

	offset := 12
	an.Sequence = data[offset]
	an.StartCode = data[offset+1]
	an.SubUni = data[offset+2]
	an.Net = data[offset+3]

	dataLength := int(binary.LittleEndian.Uint16(data[offset+4 : offset+6]))

	if len(data[offset+6:]) < dataLength {
		return errors.New("[]byte length not long enough to contain data length specified in packet")
	}
	an.Data = make([]uint8, dataLength)
	copy(an.Data, data[offset+6:offset+6+dataLength])
	return nil
}

func (an *ArtNzs) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+8)
	copy(data[0:8], an.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], an.OpCode)
	data[10] = an.ProtVerHi
	data[11] = an.ProtVerLo
	data[12] = an.Sequence
	data[13] = an.StartCode
	data[14] = an.SubUni
	data[15] = an.Net
	dataLength := uint16(len(an.Data))
	if dataLength > 512 {
		return nil, errors.New("data length must be less than or equal to 512 bytes")
	}
	binary.LittleEndian.PutUint16(data[16:18], dataLength)
	copy(data[18:], an.Data)
	return data, nil
}
