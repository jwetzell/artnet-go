package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtIpProgReply struct {
	filler1    uint8
	filler2    uint8
	filler3    uint8
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
	Status     uint8
	spare2     uint8
	ProgDgHi   uint8
	ProgDg2    uint8
	ProgDg1    uint8
	ProgDgLo   uint8
	spare7     uint8
	spare8     uint8
}

func (aipr *ArtIpProgReply) GetOpCode() uint16 {
	return OpIpProgReply
}

func (aipr *ArtIpProgReply) UnmarshalBinary(data []byte) error {
	if len(data) < 34 {
		return errors.New("ArtIpProgReply packet must be at least 18 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpIpProgReply {
		return errors.New("packet does not have the correct OpCode for an ArtIpProgReply packet")
	}

	offset := 12

	aipr.filler1 = data[offset]
	aipr.filler2 = data[offset+1]
	aipr.filler3 = data[offset+2]
	aipr.filler4 = data[offset+3]

	aipr.ProgIpHi = data[offset+4]
	aipr.ProgIp2 = data[offset+5]
	aipr.ProgIp1 = data[offset+6]
	aipr.ProgIpLo = data[offset+7]

	aipr.ProgSmHi = data[offset+8]
	aipr.ProgSm2 = data[offset+9]
	aipr.ProgSm1 = data[offset+10]
	aipr.ProgSmLo = data[offset+11]

	aipr.ProgPortHi = data[offset+12]
	aipr.ProgPortLo = data[offset+13]

	aipr.Status = data[offset+14]
	aipr.spare2 = data[offset+15]

	aipr.ProgDgHi = data[offset+16]
	aipr.ProgDg2 = data[offset+17]
	aipr.ProgDg1 = data[offset+18]
	aipr.ProgDgLo = data[offset+19]

	aipr.spare7 = data[offset+20]
	aipr.spare8 = data[offset+21]

	return nil
}

func (aipr *ArtIpProgReply) MarshalBinary() ([]byte, error) {
	data := make([]byte, 12+22)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpIpProgReply)
	data[10] = 0
	data[11] = 14
	data[12] = aipr.filler1
	data[13] = aipr.filler2
	data[14] = aipr.filler3
	data[15] = aipr.filler4
	data[16] = aipr.ProgIpHi
	data[17] = aipr.ProgIp2
	data[18] = aipr.ProgIp1
	data[19] = aipr.ProgIpLo
	data[20] = aipr.ProgSmHi
	data[21] = aipr.ProgSm2
	data[22] = aipr.ProgSm1
	data[23] = aipr.ProgSmLo
	data[24] = aipr.ProgPortHi
	data[25] = aipr.ProgPortLo
	data[26] = aipr.Status
	data[27] = aipr.spare2
	data[28] = aipr.ProgDgHi
	data[29] = aipr.ProgDg2
	data[30] = aipr.ProgDg1
	data[31] = aipr.ProgDgLo
	data[32] = aipr.spare7
	data[33] = aipr.spare8
	return data, nil
}
