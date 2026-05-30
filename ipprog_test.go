package artnet_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
			diff := cmp.Diff(test.Expected, got, cmpopts.IgnoreUnexported(artnet.ArtIpProg{}))
			if diff != "" {
				t.Fatalf("ArtIpProg does not match\n%s", diff)
			}
		})
	}
}

func BenchmarkArtIpProgUnmarshalBinary(b *testing.B) {
	data := []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0xf8, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	for b.Loop() {
		got := artnet.ArtIpProg{}

		err := got.UnmarshalBinary(data)
		if err != nil {
			b.Fatalf("failed to decode ArtIpProg: %s", err)
		}
	}
}

func BenchmarkArtIpProgMarshalBinary(b *testing.B) {
	data := artnet.ArtIpProg{
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
	}

	for b.Loop() {
		_, err := data.MarshalBinary()
		if err != nil {
			b.Fatalf("failed to encode ArtIpProg: %s", err)
		}
	}
}

func FuzzArtIpProgUnmarshalBinary(f *testing.F) {
	f.Add([]byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0xf8, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	f.Fuzz(func(t *testing.T, data []byte) {
		artIpProg := artnet.ArtIpProg{}
		artIpProg.UnmarshalBinary(data)
	})
}
