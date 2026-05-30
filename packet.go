package artnet

import (
	"encoding"
)

type ArtNetPacket interface {
	encoding.BinaryUnmarshaler
	encoding.BinaryMarshaler
	GetOpCode() uint16
	GetID() [8]uint8
}
