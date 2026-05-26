package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtIpProg struct {
	ID         [8]uint8
	OpCode     uint16
	ProtVerHi  uint8
	ProtVerLo  uint8
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

func (ad *ArtIpProg) GetOpCode() uint16 {
	return ad.OpCode
}

func (ad *ArtIpProg) GetProtVer() uint16 {
	return uint16(ad.ProtVerHi)<<8 + uint16(ad.ProtVerLo)
}

func (ad *ArtIpProg) GetID() [8]uint8 {
	return ad.ID
}

func (ad *ArtIpProg) UnmarshalBinary(data []byte) error {
	if len(data) < 34 {
		return errors.New("ArtIpProg packet must be at least 18 bytes long")
	}

	copy(ad.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ad.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ad.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ad.ProtVerHi = data[10]
	ad.ProtVerLo = data[11]

	offset := 12

	ad.filler1 = data[offset]
	ad.filler2 = data[offset+1]
	ad.Command = data[offset+2]
	ad.filler4 = data[offset+3]

	ad.ProgIpHi = data[offset+4]
	ad.ProgIp2 = data[offset+5]
	ad.ProgIp1 = data[offset+6]
	ad.ProgIpLo = data[offset+7]

	ad.ProgSmHi = data[offset+8]
	ad.ProgSm2 = data[offset+9]
	ad.ProgSm1 = data[offset+10]
	ad.ProgSmLo = data[offset+11]

	ad.ProgPortHi = data[offset+12]
	ad.ProgPortLo = data[offset+13]

	ad.ProgDgHi = data[offset+14]
	ad.ProgDg2 = data[offset+15]
	ad.ProgDg1 = data[offset+16]
	ad.ProgDgLo = data[offset+17]

	ad.spare4 = data[offset+18]
	ad.spare5 = data[offset+19]
	ad.spare6 = data[offset+20]
	ad.spare7 = data[offset+21]

	return nil
}

func (ad *ArtIpProg) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+26)
	copy(data[0:8], ad.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ad.OpCode)
	data[10] = ad.ProtVerHi
	data[11] = ad.ProtVerLo
	data[12] = ad.filler1
	data[13] = ad.filler2
	data[14] = ad.Command
	data[15] = ad.filler4
	data[16] = ad.ProgIpHi
	data[17] = ad.ProgIp2
	data[18] = ad.ProgIp1
	data[19] = ad.ProgIpLo
	data[20] = ad.ProgSmHi
	data[21] = ad.ProgSm2
	data[22] = ad.ProgSm1
	data[23] = ad.ProgSmLo
	data[24] = ad.ProgPortHi
	data[25] = ad.ProgPortLo
	data[26] = ad.ProgDgHi
	data[27] = ad.ProgDg2
	data[28] = ad.ProgDg1
	data[29] = ad.ProgDgLo
	data[30] = ad.spare4
	data[31] = ad.spare5
	data[32] = ad.spare6
	data[33] = ad.spare7
	return data, nil
}
