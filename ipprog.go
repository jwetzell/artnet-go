package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtIpProg struct {
	filler1    uint8
	filler2    uint8
	Command    uint8
	filler4    uint8
	ProgIpHi   uint8
	ProgIp2    uint8
	ProgIp1    uint8
	ProgIpLo   uint8
	ProgSmHi   uint8
	ProgSm2    uint8
	ProgSm1    uint8
	ProgSmLo   uint8
	ProgPortHi uint8
	ProgPortLo uint8
	ProgDgHi   uint8
	ProgDg2    uint8
	ProgDg1    uint8
	ProgDgLo   uint8
	spare4     uint8
	spare5     uint8
	spare6     uint8
	spare7     uint8
}

func (aip *ArtIpProg) GetOpCode() uint16 {
	return OpIpProg
}

func (aip *ArtIpProg) UnmarshalBinary(data []byte) error {
	if len(data) < 34 {
		return errors.New("ArtIpProg packet must be at least 18 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpIpProg {
		return errors.New("packet does not have the correct OpCode for an ArtIpProg packet")
	}

	offset := 12

	aip.filler1 = data[offset]
	aip.filler2 = data[offset+1]
	aip.Command = data[offset+2]
	aip.filler4 = data[offset+3]

	aip.ProgIpHi = data[offset+4]
	aip.ProgIp2 = data[offset+5]
	aip.ProgIp1 = data[offset+6]
	aip.ProgIpLo = data[offset+7]

	aip.ProgSmHi = data[offset+8]
	aip.ProgSm2 = data[offset+9]
	aip.ProgSm1 = data[offset+10]
	aip.ProgSmLo = data[offset+11]

	aip.ProgPortHi = data[offset+12]
	aip.ProgPortLo = data[offset+13]

	aip.ProgDgHi = data[offset+14]
	aip.ProgDg2 = data[offset+15]
	aip.ProgDg1 = data[offset+16]
	aip.ProgDgLo = data[offset+17]

	aip.spare4 = data[offset+18]
	aip.spare5 = data[offset+19]
	aip.spare6 = data[offset+20]
	aip.spare7 = data[offset+21]

	return nil
}

func (aip *ArtIpProg) MarshalBinary() ([]byte, error) {
	data := make([]byte, 12+22)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpIpProg)
	data[10] = 0
	data[11] = 14
	data[12] = aip.filler1
	data[13] = aip.filler2
	data[14] = aip.Command
	data[15] = aip.filler4
	data[16] = aip.ProgIpHi
	data[17] = aip.ProgIp2
	data[18] = aip.ProgIp1
	data[19] = aip.ProgIpLo
	data[20] = aip.ProgSmHi
	data[21] = aip.ProgSm2
	data[22] = aip.ProgSm1
	data[23] = aip.ProgSmLo
	data[24] = aip.ProgPortHi
	data[25] = aip.ProgPortLo
	data[26] = aip.ProgDgHi
	data[27] = aip.ProgDg2
	data[28] = aip.ProgDg1
	data[29] = aip.ProgDgLo
	data[30] = aip.spare4
	data[31] = aip.spare5
	data[32] = aip.spare6
	data[33] = aip.spare7
	return data, nil
}
