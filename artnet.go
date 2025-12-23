package artnet

import "fmt"

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

func Decode(bytes []byte) (ArtNetPacket, error) {
	header, err := NewHeader(bytes)
	if err != nil {
		return nil, err
	}
	switch header.OpCode {
	case OpPoll:
		return NewArtPoll(header, bytes)
	case OpCommand:
		return NewArtCommand(header, bytes)
	case OpDmx:
		return NewArtDmx(header, bytes)
	default:
		return nil, fmt.Errorf("unhandled opcode: %#x", header.OpCode)

	}
}
