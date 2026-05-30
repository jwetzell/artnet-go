package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type ArtPollReply struct {
	ID          [8]uint8
	OpCode      uint16
	IPAddress   [4]uint8
	Port        uint16
	VersInfo    uint16
	NetSwitch   uint8
	SubSwitch   uint8
	Oem         uint16
	UbeaVersion uint8
	Status1     uint8
	EstaMan     uint16
	PortName    [18]uint8
	LongName    [64]uint8
	NodeReport  [64]uint8
	NumPorts    uint16
	PortTypes   [4]uint8
	GoodInput   [4]uint8
	GoodOutputA [4]uint8
	SwIn        [4]uint8
	SwOut       [4]uint8
	AcnPriority uint8
	SwMacro     uint8
	SwRemote    uint8
	Spare1      uint8
	Spare2      uint8
	Spare3      uint8
	Style       uint8
	MAC         [6]uint8
	// TODO(jwetzell): support extended poll fields
	BindIp                [4]uint8
	BindIndex             uint8
	Status2               uint8
	GoodOutputB           [4]uint8
	Status3               uint8
	DefaultRespUID        uint64
	User                  uint16
	RefreshRate           uint8
	BackgroundQueuePolicy uint8
	filler                [10]uint8
}

func (ap *ArtPollReply) GetOpCode() uint16 {
	return ap.OpCode
}

func (ap *ArtPollReply) GetID() [8]uint8 {
	return ap.ID
}

func (ap *ArtPollReply) UnmarshalBinary(data []byte) error {

	if len(data) < 207 {
		return errors.New("ArtPollReply packet must be at least 207 bytes long")
	}

	copy(ap.ID[:], data[0:8])

	if !slices.Equal(ArtNetID[:], ap.ID[:]) {
		return errors.New("ID does not match Art-Net ID")
	}

	ap.OpCode = binary.LittleEndian.Uint16(data[8:10])
	ap.IPAddress[0] = data[10]
	ap.IPAddress[1] = data[11]
	ap.IPAddress[2] = data[12]
	ap.IPAddress[3] = data[13]
	ap.Port = binary.LittleEndian.Uint16(data[14:16])
	ap.VersInfo = binary.BigEndian.Uint16(data[16:18])
	ap.NetSwitch = data[18]
	ap.SubSwitch = data[19]
	ap.Oem = binary.BigEndian.Uint16(data[20:22])
	ap.UbeaVersion = data[22]
	ap.Status1 = data[23]
	ap.EstaMan = binary.LittleEndian.Uint16(data[24:26])
	copy(ap.PortName[:], data[26:44])
	copy(ap.LongName[:], data[44:108])
	copy(ap.NodeReport[:], data[108:172])
	ap.NumPorts = binary.BigEndian.Uint16(data[172:174])
	copy(ap.PortTypes[:], data[174:178])
	copy(ap.GoodInput[:], data[178:182])
	copy(ap.GoodOutputA[:], data[182:186])
	copy(ap.SwIn[:], data[186:190])
	copy(ap.SwOut[:], data[190:194])
	ap.AcnPriority = data[194]
	ap.SwMacro = data[195]
	ap.SwRemote = data[196]
	ap.Spare1 = data[197]
	ap.Spare2 = data[198]
	ap.Spare3 = data[199]
	ap.Style = data[200]
	copy(ap.MAC[:], data[201:207])
	return nil
}

func (ap *ArtPollReply) MarshalBinary() ([]byte, error) {
	if ap.DefaultRespUID > 0xFFFFFFFFFFFF {
		return nil, errors.New("DefaultRespUID must not be greater than 281474976710655")
	}
	data := make([]byte, 8+230)
	copy(data[0:8], ap.ID[:])
	binary.LittleEndian.PutUint16(data[8:10], ap.OpCode)
	data[10] = ap.IPAddress[0]
	data[11] = ap.IPAddress[1]
	data[12] = ap.IPAddress[2]
	data[13] = ap.IPAddress[3]
	binary.LittleEndian.PutUint16(data[14:16], ap.Port)
	binary.BigEndian.PutUint16(data[16:18], ap.VersInfo)
	data[18] = ap.NetSwitch
	data[19] = ap.SubSwitch
	binary.BigEndian.PutUint16(data[20:22], ap.Oem)
	data[22] = ap.UbeaVersion
	data[23] = ap.Status1
	binary.LittleEndian.PutUint16(data[24:26], ap.EstaMan)
	copy(data[26:44], ap.PortName[:])
	copy(data[44:108], ap.LongName[:])
	copy(data[108:172], ap.NodeReport[:])
	binary.BigEndian.PutUint16(data[172:174], ap.NumPorts)
	copy(data[174:178], ap.PortTypes[:])
	copy(data[178:182], ap.GoodInput[:])
	copy(data[182:186], ap.GoodOutputA[:])
	copy(data[186:190], ap.SwIn[:])
	copy(data[190:194], ap.SwOut[:])
	data[194] = ap.AcnPriority
	data[195] = ap.SwMacro
	data[196] = ap.SwRemote
	data[197] = ap.Spare1
	data[198] = ap.Spare2
	data[199] = ap.Spare3
	data[200] = ap.Style
	copy(data[201:207], ap.MAC[:])
	copy(data[207:211], ap.BindIp[:])
	data[211] = ap.BindIndex
	data[212] = ap.Status2
	copy(data[213:217], ap.GoodOutputB[:])
	data[217] = ap.Status3
	// DefaultRespUID is a 48-bit value, so we need to make sure we only write the lower 6 bytes
	data[218] = uint8(ap.DefaultRespUID >> 40)
	data[219] = uint8(ap.DefaultRespUID >> 32)
	data[220] = uint8(ap.DefaultRespUID >> 24)
	data[221] = uint8(ap.DefaultRespUID >> 16)
	data[222] = uint8(ap.DefaultRespUID >> 8)
	data[223] = uint8(ap.DefaultRespUID)
	binary.BigEndian.PutUint16(data[224:226], ap.User)
	data[226] = ap.RefreshRate
	data[227] = ap.BackgroundQueuePolicy
	copy(data[228:], ap.filler[:])
	return data, nil
}
