package artnet

import (
	"errors"
)

type ArtDmx struct {
	Header   *ArtNetHeader
	Sequence uint8
	Physical uint8
	SubUni   uint8
	Net      uint8
	Length   uint16
	Data     []uint8
}

func NewArtDmx(header *ArtNetHeader, data []byte) (*ArtDmx, error) {
	artDmx := ArtDmx{
		Header: header,
	}

	if len(data) < 18 {
		return nil, errors.New("ArtDmx packet must be at least 14 bytes long")
	}
	offset := 12

	artDmx.Sequence = data[offset]
	artDmx.Physical = data[offset+1]
	artDmx.SubUni = data[offset+2]
	artDmx.Net = data[offset+3]

	artDmx.Length = uint16(data[offset+4])<<8 + uint16(data[offset+5])

	dmxDataOffset := offset + 6

	if len(data[dmxDataOffset:]) < int(artDmx.Length) {
		return nil, errors.New("ArtDmx packet length mismatch")
	}

	artDmx.Data = make([]uint8, artDmx.Length)

	copy(artDmx.Data, data[dmxDataOffset:dmxDataOffset+int(artDmx.Length)])

	return &artDmx, nil
}

func (ap *ArtDmx) GetOpCode() uint16 {
	return ap.Header.OpCode
}
