package artnet_test

import (
	"reflect"
	"testing"

	"github.com/jwetzell/artnet-go"
)

func TestGoodArtIpProgUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtIpProg
	}{
		{
			Name: "ACT Packet 1",
			Data: []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0xf8, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			Expected: &artnet.ArtIpProg{
				Command:    0x00,
				ProgIpHi:   0x00,
				ProgIp2:    0x00,
				ProgIp1:    0x00,
				ProgIpLo:   0x00,
				ProgSmHi:   0x00,
				ProgSm2:    0x00,
				ProgSm1:    0x00,
				ProgSmLo:   0x00,
				ProgPortHi: 0x00,
				ProgPortLo: 0x00,
				ProgDgHi:   0x00,
				ProgDg2:    0x00,
				ProgDg1:    0x00,
				ProgDgLo:   0x00,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtIpProg{}

			err := got.UnmarshalBinary(test.Data)
			if err != nil {
				t.Fatalf("failed to Unmarshal ArtIpProg: %s", err)
			}
			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("ArtIpProg does not match got: %+v expected: %+v", got, test.Expected)
			}
		})
	}
}
