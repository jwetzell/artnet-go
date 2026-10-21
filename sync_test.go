package artnet_test

import (
	"testing"

	"github.com/jwetzell/artnet-go"
	"github.com/stretchr/testify/require"
)

func TestGoodArtSyncUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtSync
	}{
		{
			Name:     "zero",
			Data:     []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x52, 0x00, 0x0e, 0x00, 0x00},
			Expected: &artnet.ArtSync{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtSync{}

			err := got.UnmarshalBinary(test.Data)
			require.NoError(t, err)

			require.Equal(t, artnet.OpSync, got.GetOpCode())

			require.Equal(t, test.Expected, got)
		})
	}
}

func TestBadArtSyncUnmarshal(t *testing.T) {
	tests := []struct {
		Name          string
		Data          []byte
		ExpectedError string
	}{
		{
			Name:          "Too Short",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x52, 0x00, 0x0e},
			ExpectedError: "ArtSync packet must be at least 14 bytes long",
		},
		{
			Name:          "Invalid Header",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x73, 0x00, 0x00, 0x52, 0x00, 0x0e, 0x00, 0x00},
			ExpectedError: "ID does not match Art-Net ID",
		},
		{
			Name:          "Incorrect OpCode",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x51, 0x00, 0x0e, 0x00, 0x00},
			ExpectedError: "packet does not have the correct OpCode for an ArtSync packet",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtSync{}
			err := got.UnmarshalBinary(test.Data)
			require.Error(t, err)
			require.EqualError(t, err, test.ExpectedError)
		})
	}
}

func TestGoodArtSyncMarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Expected []byte
		Data     *artnet.ArtSync
	}{
		{
			Name:     "zero",
			Expected: []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x52, 0x00, 0x0e, 0x00, 0x00},
			Data:     &artnet.ArtSync{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got, err := test.Data.MarshalBinary()
			require.NoError(t, err)
			require.Equal(t, test.Expected, got)
		})
	}
}

func TestBadArtSyncMarshal(t *testing.T) {
	tests := []struct {
		Name          string
		Data          *artnet.ArtSync
		ExpectedError string
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := test.Data.MarshalBinary()
			require.Error(t, err)
			require.EqualError(t, err, test.ExpectedError)
		})
	}
}

func BenchmarkArtSyncMarshalBinary(b *testing.B) {
	data := artnet.ArtSync{}

	for b.Loop() {
		_, err := data.MarshalBinary()
		require.NoError(b, err)
	}
}
