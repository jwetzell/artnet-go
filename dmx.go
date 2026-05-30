package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtDmx struct {
	Sequence uint8
	Physical uint8
	SubUni   uint8
	Net      uint8
	Data     []uint8
}

func (ad *ArtDmx) GetOpCode() uint16 {
	return OpDmx
}

func (ad *ArtDmx) UnmarshalBinary(data []byte) error {
	if len(data) < 18 {
		return errors.New("ArtDmx packet must be at least 18 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpDmx {
		return errors.New("packet does not have the correct OpCode for an ArtDmx packet")
	}

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
	// TODO(jwetzell): check max data length
	data := make([]byte, 12+6+len(ad.Data))
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpDmx)
	data[10] = 0
	data[11] = 14
	data[12] = ad.Sequence
	data[13] = ad.Physical
	data[14] = ad.SubUni
	data[15] = ad.Net
	binary.BigEndian.PutUint16(data[16:18], uint16(len(ad.Data)))
	copy(data[18:], ad.Data)
	return data, nil
}
