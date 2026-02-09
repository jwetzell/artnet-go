package artnet_test

import (
	"slices"
	"testing"

	"github.com/jwetzell/artnet-go"
)

func TestGoodArtTimeCode(t *testing.T) {
	tests := []struct {
		Data     []byte
		Expected *artnet.ArtTimeCode
	}{
		{
			Data: []byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 151, 0, 14, 0, 0, 11, 17, 3, 0, 0},
			Expected: &artnet.ArtTimeCode{
				ID:        []byte{'A', 'r', 't', '-', 'N', 'e', 't', 0x00},
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
		got := artnet.ArtTimeCode{}

		err := got.UnmarshalBinary(test.Data)

		if err != nil {
			t.Fatalf("failed to decode ArtTimeCode: %s", err)
		}

		if got.OpCode != test.Expected.OpCode {
			t.Fatalf("ArtTimeCode OpCode does not match got: %d expected: %d", got.OpCode, test.Expected.OpCode)
		}

		if !slices.Equal(got.ID, test.Expected.ID) {
			t.Fatalf("ArtTimeCode ID does not match got: %+v expected: %+v", got.ID, test.Expected.ID)
		}

		if got.Filler1 != test.Expected.Filler1 {
			t.Fatalf("ArtTimeCode Filler1 does not match got: %d expected: %d", got.Filler1, test.Expected.Filler1)
		}

		if got.StreamId != test.Expected.StreamId {
			t.Fatalf("ArtTimeCode StreamId does not match got: %d expected: %d", got.StreamId, test.Expected.StreamId)
		}

		if got.Frames != test.Expected.Frames {
			t.Fatalf("ArtTimeCode Frames does not match got: %d expected: %d", got.Frames, test.Expected.Frames)
		}

		if got.Seconds != test.Expected.Seconds {
			t.Fatalf("ArtTimeCode Seconds does not match got: %d expected: %d", got.Seconds, test.Expected.Seconds)
		}

		if got.Minutes != test.Expected.Minutes {
			t.Fatalf("ArtTimeCode Minutes does not match got: %d expected: %d", got.Minutes, test.Expected.Minutes)
		}

		if got.Hours != test.Expected.Hours {
			t.Fatalf("ArtTimeCode Hours does not match got: %d expected: %d", got.Hours, test.Expected.Hours)
		}

		if got.Type != test.Expected.Type {
			t.Fatalf("ArtTimeCode Type does not match got: %d expected: %d", got.Type, test.Expected.Type)
		}

		if got.ProtVerHi != test.Expected.ProtVerHi {
			t.Fatalf("ArtTimeCode ProtVerHi does not match got: %d expected: %d", got.ProtVerHi, test.Expected.ProtVerHi)
		}

		if got.ProtVerLo != test.Expected.ProtVerLo {
			t.Fatalf("ArtTimeCode ProtVerLo does not match got: %d expected: %d", got.ProtVerLo, test.Expected.ProtVerLo)
		}

	}
}
