package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtDmx struct {
	ID        [8]uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	Sequence  uint8
	Physical  uint8
	SubUni    uint8
	Net       uint8
	Data      []uint8
}

func (ad *ArtDmx) GetOpCode() uint16 {
	return ad.OpCode
}

func (ad *ArtDmx) GetProtVer() uint16 {
	return uint16(ad.ProtVerHi)<<8 + uint16(ad.ProtVerLo)
}

func (ad *ArtDmx) GetID() [8]uint8 {
	return ad.ID
}

func (ad *ArtDmx) Length() uint16 {
	return uint16(len(ad.Data))
}

func (ad *ArtDmx) UnmarshalBinary(data []byte) error {
	if len(data) < 18 {
		return errors.New("ArtDmx packet must be at least 18 bytes long")
	}

	copy(ad.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ad.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ad.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ad.ProtVerHi = data[10]
	ad.ProtVerLo = data[11]

	offset := 12

	ad.Sequence = data[offset]
	ad.Physical = data[offset+1]
	ad.SubUni = data[offset+2]
	ad.Net = data[offset+3]

	length := int(data[offset+4])<<8 + int(data[offset+5])

	dmxDataOffset := offset + 6

	if len(data[dmxDataOffset:]) < length {
		return errors.New("ArtDmx packet length mismatch")
	}

	ad.Data = data[dmxDataOffset : dmxDataOffset+length]

	return nil
}

func (ad *ArtDmx) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+10+ad.Length())
	copy(data[0:8], ad.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ad.OpCode)
	data[10] = ad.ProtVerHi
	data[11] = ad.ProtVerLo
	data[12] = ad.Sequence
	data[13] = ad.Physical
	data[14] = ad.SubUni
	data[15] = ad.Net
	binary.BigEndian.PutUint16(data[16:18], ad.Length())
	copy(data[18:], ad.Data)
	return data, nil
}
