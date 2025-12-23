package artnet

import (
	"errors"
)

type ArtPoll struct {
	Header          *ArtNetHeader
	Flags           uint8
	DiagPriority    uint8
	AddressTopHi    uint8
	AddressTopLo    uint8
	AddressBottomHi uint8
	AddressBottomLo uint8
	EstaManHi       uint8
	EstaManLo       uint8
	OemHi           uint8
	OemLo           uint8
}

func NewArtPoll(header *ArtNetHeader, data []byte) (*ArtPoll, error) {
	artPoll := ArtPoll{
		Header: header,
	}

	if len(data) < 14 {
		return nil, errors.New("ArtPoll packet must be at least 14 bytes long")
	}
	offset := 12
	artPoll.Flags = data[offset]
	artPoll.DiagPriority = data[offset+1]

	//TODO(jwetzell): unpack extended poll fields

	return &artPoll, nil
}

func (ap *ArtPoll) GetOpCode() uint16 {
	return ap.Header.OpCode
}
