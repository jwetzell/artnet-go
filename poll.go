package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtPoll struct {
	ID           [8]uint8
	OpCode       uint16
	ProtVerHi    uint8
	ProtVerLo    uint8
	Flags        uint8
	DiagPriority uint8
	// TODO(jwetzell): support extended poll fields
	TargetPortAddressTop    uint16
	TargetPortAddressBottom uint16
	EstaMan                 uint16
	Oem                     uint16
}

func (ap *ArtPoll) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtPoll) GetProtVer() uint16 {
	return uint16(ap.ProtVerHi)<<8 + uint16(ap.ProtVerLo)
}

func (ap *ArtPoll) GetID() [8]uint8 {
	return ap.ID
}

func (ap *ArtPoll) UnmarshalBinary(data []byte) error {

	if len(data) < 14 {
		return errors.New("ArtPoll packet must be at least 14 bytes long")
	}

	copy(ap.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ap.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ap.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ap.ProtVerHi = data[10]
	ap.ProtVerLo = data[11]

	offset := 12
	ap.Flags = data[offset]
	ap.DiagPriority = data[offset+1]

	if len(data) >= 16 {
		ap.TargetPortAddressTop = binary.BigEndian.Uint16(data[14:16])
	}
	if len(data) >= 18 {
		ap.TargetPortAddressBottom = binary.BigEndian.Uint16(data[16:18])
	}
	if len(data) >= 20 {
		ap.EstaMan = binary.BigEndian.Uint16(data[18:20])
	}
	if len(data) >= 22 {
		ap.Oem = binary.BigEndian.Uint16(data[20:22])
	}
	return nil
}

func (ap *ArtPoll) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+14)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.ProtVerHi
	data[11] = ap.ProtVerLo
	data[12] = ap.Flags
	data[13] = ap.DiagPriority
	binary.BigEndian.PutUint16(data[14:16], ap.TargetPortAddressTop)
	binary.BigEndian.PutUint16(data[16:18], ap.TargetPortAddressBottom)
	binary.BigEndian.PutUint16(data[18:20], ap.EstaMan)
	binary.BigEndian.PutUint16(data[20:22], ap.Oem)
	return data, nil
}
