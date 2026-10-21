package artnet_test

import (
	"testing"

	"github.com/jwetzell/artnet-go"
	"github.com/stretchr/testify/require"
)

func TestGoodArtTodControlUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtTodControl
	}{
		{
			Name: "ACT Packet 1",
			Data: []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x82, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00},
			Expected: &artnet.ArtTodControl{
				Net:     0x00,
				Command: artnet.AtcFlush,
				Address: 0x00,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtTodControl{}

			err := got.UnmarshalBinary(test.Data)
			require.NoError(t, err)

			require.Equal(t, artnet.OpTodControl, got.GetOpCode())

			require.Equal(t, test.Expected, got)
		})
	}
}

func BenchmarkArtTodControlUnmarshalBinary(b *testing.B) {
	data := []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x82, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00}

	for b.Loop() {
		got := artnet.ArtTodControl{}

		err := got.UnmarshalBinary(data)
		require.NoError(b, err)
	}
}

func BenchmarkArtTodControlMarshalBinary(b *testing.B) {
	data := artnet.ArtTodControl{
		Net:     0x00,
		Command: artnet.AtcFlush,
		Address: 0x00,
	}

	for b.Loop() {
		_, err := data.MarshalBinary()
		require.NoError(b, err)
	}
}

func FuzzArtTodControlUnmarshalBinary(f *testing.F) {
	f.Add([]byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x82, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00})
	f.Fuzz(func(t *testing.T, data []byte) {
		artTodControl := artnet.ArtTodControl{}
		artTodControl.UnmarshalBinary(data)
	})
}
