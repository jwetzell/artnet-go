package artnet

import (
	"errors"
)

type ArtCommand struct {
	Header    *ArtNetHeader
	EstaManHi uint8
	EstaManLo uint8
	Length    uint16
	Data      string
}

func NewArtCommand(header *ArtNetHeader, data []byte) (*ArtCommand, error) {
	artCommand := ArtCommand{
		Header: header,
	}

	if len(data) < 18 {
		return nil, errors.New("ArtCommand packet must be at least 14 bytes long")
	}
	offset := 12

	artCommand.EstaManHi = data[offset]
	artCommand.EstaManLo = data[offset+1]

	artCommand.Length = uint16(data[offset+2])<<8 + uint16(data[offset+3])

	commandDataOffset := offset + 4

	if len(data[commandDataOffset:]) < int(artCommand.Length) {
		return nil, errors.New("ArtCommand packet length mismatch")
	}

	artCommand.Data = string(data[commandDataOffset : commandDataOffset+int(artCommand.Length)-1])

	return &artCommand, nil
}

func (ap *ArtCommand) GetOpCode() uint16 {
	return ap.Header.OpCode
}
