package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtPoll struct {
	ID              []uint8
	OpCode          uint16
	ProtVerHi       uint8
	ProtVerLo       uint8
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

func NewArtPoll(data []byte) (*ArtPoll, error) {
	artPoll := ArtPoll{}

	err := artPoll.UnmarshalBinary(data)

	if err != nil {
		return nil, err
	}

	return &artPoll, nil
}

func (ap *ArtPoll) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtPoll) GetProtVer() uint16 {
	return uint16(ap.ProtVerHi)<<8 + uint16(ap.ProtVerLo)
}

func (ap *ArtPoll) GetID() []uint8 {
	return ap.ID
}

func (ap *ArtPoll) UnmarshalBinary(data []byte) error {

	if len(data) < 14 {
		return errors.New("ArtPoll packet must be at least 14 bytes long")
	}

	ap.ID = data[0:8]

	if !slices.Equal(ArtNetID, ap.ID) {
		return errors.New("ID does not match Art-Net ID")
	}

	ap.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ap.ProtVerHi = data[10]
	ap.ProtVerLo = data[11]

	offset := 12
	ap.Flags = data[offset]
	ap.DiagPriority = data[offset+1]
	//TODO(jwetzell): unpack extended poll fields
	return nil
}

func (ap *ArtPoll) MarshalBinary() ([]byte, error) {
	data := []byte(ArtNetID)
	data = append(data, byte(ap.OpCode), byte(ap.OpCode>>8), ap.ProtVerHi, ap.ProtVerLo, ap.Flags, ap.DiagPriority)
	//TODO(jwetzell): pack extended poll fields
	return data, nil
}
