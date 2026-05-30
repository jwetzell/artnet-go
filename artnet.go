package artnet

import (
	"encoding/binary"
	"errors"
	"fmt"
)

var (
	OpPoll             uint16 = 0x2000
	OpPollReply        uint16 = 0x2100
	OpDiagData         uint16 = 0x2300
	OpCommand          uint16 = 0x2400
	OpDataRequest      uint16 = 0x2700
	OpDataReply        uint16 = 0x2800
	OpOutput           uint16 = 0x5000
	OpDmx              uint16 = 0x5000
	OpNzs              uint16 = 0x5100
	OpSync             uint16 = 0x5200
	OpAddress          uint16 = 0x6000
	OpInput            uint16 = 0x7000
	OpTodRequest       uint16 = 0x8000
	OpTodData          uint16 = 0x8100
	OpTodControl       uint16 = 0x8200
	OpRdm              uint16 = 0x8300
	OpRdmSub           uint16 = 0x8400
	OpVideoSetup       uint16 = 0xa010
	OpVideoPalette     uint16 = 0xa020
	OpVideoData        uint16 = 0xa040
	OpMacMaster        uint16 = 0xf000
	OpMacSlave         uint16 = 0xf100
	OpFirmwareMaster   uint16 = 0xf200
	OpFirmwareReply    uint16 = 0xf300
	OpFileTnMaster     uint16 = 0xf400
	OpFileFnMaster     uint16 = 0xf500
	OpFileFnReply      uint16 = 0xf600
	OpIpProg           uint16 = 0xf800
	OpIpProgReply      uint16 = 0xf900
	OpMedia            uint16 = 0x9000
	OpMediaPatch       uint16 = 0x9100
	OpMediaControl     uint16 = 0x9200
	OpMediaContrlReply uint16 = 0x9300
	OpTimeCode         uint16 = 0x9700
	OpTimeSync         uint16 = 0x9800
	OpTrigger          uint16 = 0x9900
	OpDirectory        uint16 = 0x9a00
	OpDirectoryReply   uint16 = 0x9b00
)

var ArtNetID [8]uint8 = [8]uint8{'A', 'r', 't', '-', 'N', 'e', 't', 0x00}

func Decode(bytes []byte) (ArtNetPacket, error) {
	if len(bytes) < 12 {
		return nil, errors.New("ArtNet packet must be at least 12 bytes")
	}
	opCode := binary.LittleEndian.Uint16(bytes[8:10])
	switch opCode {
	case OpPoll:
		artPoll := ArtPoll{}

		err := artPoll.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artPoll, nil
	case OpIpProg:
		artIpProg := ArtIpProg{}

		err := artIpProg.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artIpProg, nil
	case OpIpProgReply:
		artIpProgReply := ArtIpProgReply{}

		err := artIpProgReply.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artIpProgReply, nil
	case OpDmx:
		artDmx := ArtDmx{}

		err := artDmx.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artDmx, nil
	case OpTimeCode:
		artTimeCode := ArtTimeCode{}

		err := artTimeCode.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artTimeCode, nil
	case OpDataRequest:
		artDataRequest := ArtDataRequest{}

		err := artDataRequest.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artDataRequest, nil
	case OpDiagData:
		artDiagData := ArtDiagData{}

		err := artDiagData.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artDiagData, nil
	case OpCommand:
		artCommand := ArtCommand{}

		err := artCommand.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artCommand, nil
	case OpTrigger:
		artTrigger := ArtTrigger{}

		err := artTrigger.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artTrigger, nil
	case OpSync:
		artSync := ArtSync{}

		err := artSync.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artSync, nil
	case OpTodRequest:
		artTodRequest := ArtTodRequest{}

		err := artTodRequest.UnmarshalBinary(bytes)

		if err != nil {
			return nil, err
		}

		return &artTodRequest, nil
	default:
		return nil, fmt.Errorf("unhandled opcode: %#x", opCode)
	}
}
