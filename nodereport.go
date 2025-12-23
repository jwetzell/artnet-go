package artnet

var (
	RcDebug        uint16 = 0x0000
	RcPowerOk      uint16 = 0x0001
	RcPowerFail    uint16 = 0x0002
	RcSocketWr1    uint16 = 0x0003
	RcParseFail    uint16 = 0x0004
	RcUdpFail      uint16 = 0x0005
	RcShNameOk     uint16 = 0x0006
	RcLoNameOk     uint16 = 0x0007
	RcDmxError     uint16 = 0x0008
	RcDmxUdpFull   uint16 = 0x0009
	RcDmxRxFull    uint16 = 0x000a
	RcSwitchErr    uint16 = 0x000b
	RcConfigErr    uint16 = 0x000c
	RcDmxShort     uint16 = 0x000d
	RcFirmwareFail uint16 = 0x000e
	RcUserFail     uint16 = 0x000f
	RcFactoryRes   uint16 = 0x0010
)
