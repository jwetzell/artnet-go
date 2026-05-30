package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtNzs struct {
	Sequence  uint8
	StartCode uint8
	SubUni    uint8
	Net       uint8
	Data      []uint8
}

func (an *ArtNzs) GetOpCode() uint16 {
	return OpNzs
}

func (an *ArtNzs) UnmarshalBinary(data []byte) error {

	if len(data) < 18 {
		return errors.New("ArtNzs packet must be at least 18 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpNzs {
		return errors.New("packet does not have the correct OpCode for an ArtNzs packet")
	}

	offset := 12
	an.Sequence = data[offset]
	an.StartCode = data[offset+1]
	an.SubUni = data[offset+2]
	an.Net = data[offset+3]

	dataLength := int(binary.BigEndian.Uint16(data[offset+4 : offset+6]))

	if len(data[offset+6:]) < dataLength {
		return errors.New("[]byte length not long enough to contain data length specified in packet")
	}
	an.Data = make([]uint8, dataLength)
	copy(an.Data, data[offset+6:offset+6+dataLength])
	return nil
}

func (an *ArtNzs) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+8)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpNzs)
	data[10] = 0
	data[11] = 14
	data[12] = an.Sequence
	data[13] = an.StartCode
	data[14] = an.SubUni
	data[15] = an.Net
	dataLength := uint16(len(an.Data))
	if dataLength > 512 {
		return nil, errors.New("data length must be less than or equal to 512 bytes")
	}
	binary.BigEndian.PutUint16(data[16:18], dataLength)
	copy(data[18:], an.Data)
	return data, nil
}
