package artnet

import (
	"encoding/binary"
	"errors"
	"slices"
)

type AddressCommand uint8

const (
	AcNone         AddressCommand = 0x00
	AcCancelMerge  AddressCommand = 0x01
	AcLedNormal    AddressCommand = 0x02
	AcLedMute      AddressCommand = 0x03
	AcLedLocate    AddressCommand = 0x04
	AcResetRxFlags AddressCommand = 0x05
	AcAnalysisOn   AddressCommand = 0x06
	AcAnalysisOff  AddressCommand = 0x07
	AcFailHold     AddressCommand = 0x08
	AcFailZero     AddressCommand = 0x09
	AcFailFull     AddressCommand = 0x0A
	AcFailScene    AddressCommand = 0x0B
	AcFailRecord   AddressCommand = 0x0C
	AcMergeLtp0    AddressCommand = 0x10
	AcMergeLtp1    AddressCommand = 0x11
	AcMergeLtp2    AddressCommand = 0x12
	AcMergeLtp3    AddressCommand = 0x13
	AcDirectionTx0 AddressCommand = 0x20
	AcDirectionTx1 AddressCommand = 0x21
	AcDirectionTx2 AddressCommand = 0x22
	AcDirectionTx3 AddressCommand = 0x23
	AcDirectionRx0 AddressCommand = 0x30
	AcDirectionRx1 AddressCommand = 0x31
	AcDirectionRx2 AddressCommand = 0x32
	AcDirectionRx3 AddressCommand = 0x33
	AcMergeHtp0    AddressCommand = 0x50
	AcMergeHtp1    AddressCommand = 0x51
	AcMergeHtp2    AddressCommand = 0x52
	AcMergeHtp3    AddressCommand = 0x53
	AcArtNetSel0   AddressCommand = 0x60
	AcArtNetSel1   AddressCommand = 0x61
	AcArtNetSel2   AddressCommand = 0x62
	AcArtNetSel3   AddressCommand = 0x63
	AcAcnSel0      AddressCommand = 0x70
	AcAcnSel1      AddressCommand = 0x71
	AcAcnSel2      AddressCommand = 0x72
	AcAcnSel3      AddressCommand = 0x73
	AcClearOp0     AddressCommand = 0x90
	AcClearOp1     AddressCommand = 0x91
	AcClearOp2     AddressCommand = 0x92
	AcClearOp3     AddressCommand = 0x93
	AcStyleDelta0  AddressCommand = 0xa0
	AcStyleDelta1  AddressCommand = 0xa1
	AcStyleDelta2  AddressCommand = 0xa2
	AcStyleDelta3  AddressCommand = 0xa3
	AcStyleConst0  AddressCommand = 0xb0
	AcStyleConst1  AddressCommand = 0xb1
	AcStyleConst2  AddressCommand = 0xb2
	AcStyleConst3  AddressCommand = 0xb3
	AcRdmEnable0   AddressCommand = 0xc0
	AcRdmEnable1   AddressCommand = 0xc1
	AcRdmEnable2   AddressCommand = 0xc2
	AcRdmEnable3   AddressCommand = 0xc3
	AcRdmDisable0  AddressCommand = 0xd0
	AcRdmDisable1  AddressCommand = 0xd1
	AcRdmDisable2  AddressCommand = 0xd2
	AcRdmDisable3  AddressCommand = 0xd3
	AcBqp0         AddressCommand = 0xe0
	AcBqp1         AddressCommand = 0xe1
	AcBqp2         AddressCommand = 0xe2
	AcBqp3         AddressCommand = 0xe3
	AcBqp4         AddressCommand = 0xe4
	AcBqp5         AddressCommand = 0xe5
	AcBqp6         AddressCommand = 0xe6
	AcBqp7         AddressCommand = 0xe7
	AcBqp8         AddressCommand = 0xe8
	AcBqp9         AddressCommand = 0xe9
	AcBqp10        AddressCommand = 0xea
	AcBqp11        AddressCommand = 0xeb
	AcBqp12        AddressCommand = 0xec
	AcBqp13        AddressCommand = 0xed
	AcBqp14        AddressCommand = 0xee
	AcBqp15        AddressCommand = 0xef
)

type ArtAddress struct {
	NetSwitch   uint8
	BindIndex   uint8
	PortName    [18]uint8
	LongName    [64]uint8
	SwIn        [4]uint8
	SwOut       [4]uint8
	SubSwitch   uint8
	AcnPriority uint8
	Command     AddressCommand
}

func (aa *ArtAddress) GetOpCode() uint16 {
	return OpAddress
}

func (aa *ArtAddress) UnmarshalBinary(data []byte) error {
	if len(data) < 107 {
		return errors.New("ArtAddress packet must be at least 107 bytes long")
	}

	if !slices.Equal(ArtNetID[:], data[0:8]) {
		return errors.New("ID does not match Art-Net ID")
	}

	opCode := binary.LittleEndian.Uint16(data[8:10])
	if opCode != OpAddress {
		return errors.New("packet does not have the correct OpCode for an ArtAddress packet")
	}
	offset := 12
	aa.NetSwitch = data[offset]
	aa.BindIndex = data[offset+1]
	copy(aa.PortName[:], data[offset+2:offset+20])
	copy(aa.LongName[:], data[offset+20:offset+84])
	copy(aa.SwIn[:], data[offset+84:offset+88])
	copy(aa.SwOut[:], data[offset+88:offset+92])
	aa.SubSwitch = data[offset+92]
	aa.AcnPriority = data[offset+93]
	aa.Command = AddressCommand(data[offset+94])
	return nil
}

func (aa *ArtAddress) MarshalBinary() ([]byte, error) {
	data := make([]byte, 12+95)
	copy(data[0:8], ArtNetID[:])
	binary.LittleEndian.PutUint16(data[8:10], OpAddress)
	data[10] = 0
	data[11] = 14
	offset := 12
	data[offset] = aa.NetSwitch
	data[offset+1] = aa.BindIndex
	copy(data[offset+2:offset+20], aa.PortName[:])
	copy(data[offset+20:offset+84], aa.LongName[:])
	copy(data[offset+84:offset+88], aa.SwIn[:])
	copy(data[offset+88:offset+92], aa.SwOut[:])
	data[offset+92] = aa.SubSwitch
	data[offset+93] = aa.AcnPriority
	data[offset+94] = uint8(aa.Command)
	return data, nil
}
