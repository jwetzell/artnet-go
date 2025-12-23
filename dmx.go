package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtDmx struct {
	ID        []uint8
	OpCode    uint16
	ProtVerHi uint8
	ProtVerLo uint8
	Sequence  uint8
	Physical  uint8
	SubUni    uint8
	Net       uint8
	Length    uint16
	Data      []uint8
}

func NewArtDmx(data []byte) (*ArtDmx, error) {
	artDmx := ArtDmx{}

	err := artDmx.UnmarshalBinary(data)

	if err != nil {
		return nil, err
	}

	return &artDmx, nil
}

func (ad *ArtDmx) GetOpCode() uint16 {
	return ad.OpCode
}

func (ad *ArtDmx) GetProtVer() uint16 {
	return uint16(ad.ProtVerHi)<<8 + uint16(ad.ProtVerLo)
}

func (ad *ArtDmx) GetID() []uint8 {
	return ad.ID
}

func (ad *ArtDmx) UnmarshalBinary(data []byte) error {
	if len(data) < 18 {
		return errors.New("ArtDmx packet must be at least 18 bytes long")
	}

	if !slices.Equal(ArtNetID, data[0:8]) {
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

	ad.Length = uint16(data[offset+4])<<8 + uint16(data[offset+5])

	dmxDataOffset := offset + 6

	if len(data[dmxDataOffset:]) < int(ad.Length) {
		return errors.New("ArtDmx packet length mismatch")
	}

	ad.Data = make([]uint8, ad.Length)

	copy(ad.Data, data[dmxDataOffset:dmxDataOffset+int(ad.Length)])
	return nil
}

func (ad *ArtDmx) MarshalBinary() ([]byte, error) {
	data := []byte(ArtNetID)
	data = append(data, byte(ad.OpCode), byte(ad.OpCode>>8), ad.ProtVerHi, ad.ProtVerLo, ad.Sequence, ad.Physical, ad.SubUni, ad.Net, byte(ad.Length>>8), byte(ad.Length))
	data = append(data, ad.Data...)
	return data, nil
}
