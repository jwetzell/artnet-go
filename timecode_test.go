package artnet_test

import (
	"reflect"
	"testing"

	"github.com/jwetzell/artnet-go"
)

func TestGoodArtTimeCodeUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtTimeCode
	}{
		{
			Name: "Basic timecode",
			Data: []byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 151, 0, 14, 0, 0, 11, 17, 3, 0, 0},
			Expected: &artnet.ArtTimeCode{
				ID:        [8]uint8{'A', 'r', 't', '-', 'N', 'e', 't', 0x00},
				OpCode:    artnet.OpTimeCode,
				ProtVerHi: 0,
				ProtVerLo: 14,
				Filler1:   0,
				StreamId:  0,
				Frames:    11,
				Seconds:   17,
				Minutes:   3,
				Hours:     0,
				Type:      0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtTimeCode{}
			err := got.UnmarshalBinary(test.Data)

			if err != nil {
				t.Fatalf("failed to decode ArtTimeCode: %s", err)
			}

			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("ArtTimeCode does not match got: %+v expected: %+v", got, test.Expected)
			}
		})
	}
}

func BenchmarkArtTimeCodeUnmarshalBinary(b *testing.B) {
	data := []byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 151, 0, 14, 0, 0, 11, 17, 3, 0, 0}

	for b.Loop() {
		got := artnet.ArtTimeCode{}

		err := got.UnmarshalBinary(data)
		if err != nil {
			b.Fatalf("failed to decode ArtTimeCode: %s", err)
		}
	}
}

func BenchmarkArtTimeCodeMarshalBinary(b *testing.B) {
	data := artnet.ArtTimeCode{
		ID:        [8]uint8{'A', 'r', 't', '-', 'N', 'e', 't', 0x00},
		OpCode:    artnet.OpTimeCode,
		ProtVerHi: 0,
		ProtVerLo: 14,
		Filler1:   0,
		StreamId:  0,
		Frames:    11,
		Seconds:   17,
		Minutes:   3,
		Hours:     0,
		Type:      0,
	}

	for b.Loop() {
		_, err := data.MarshalBinary()
		if err != nil {
			b.Fatalf("failed to encode ArtTimeCode: %s", err)
		}
	}
}

func FuzzArtTimeCodeUnmarshalBinary(f *testing.F) {
	f.Add([]byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 151, 0, 14, 0, 0, 11, 17, 3, 0, 0})
	f.Fuzz(func(t *testing.T, data []byte) {
		artTimeCode := artnet.ArtTimeCode{}
		artTimeCode.UnmarshalBinary(data)
	})
}
