package artnet_test

import (
	"testing"

	"github.com/jwetzell/artnet-go"
	"github.com/stretchr/testify/require"
)

func TestGoodArtCommandUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtCommand
	}{
		{
			Name:     "zero",
			Data:     []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x24, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00},
			Expected: &artnet.ArtCommand{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtCommand{}

			err := got.UnmarshalBinary(test.Data)
			require.NoError(t, err)

			require.Equal(t, artnet.OpCommand, got.GetOpCode())

			require.Equal(t, test.Expected, got)
		})
	}
}

func TestBadArtCommandUnmarshal(t *testing.T) {
	tests := []struct {
		Name          string
		Data          []byte
		ExpectedError string
	}{
		{
			Name:          "Too Short",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x24, 0x00, 0x0e},
			ExpectedError: "ArtCommand packet must be at least 16 bytes long",
		},
		{
			Name:          "Invalid Header",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x73, 0x00, 0x00, 0x24, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00},
			ExpectedError: "ID does not match Art-Net ID",
		},
		{
			Name:          "Incorrect OpCode",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x23, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00},
			ExpectedError: "packet does not have the correct OpCode for an ArtCommand packet",
		},
		{
			Name:          "not enough bytes for Data size",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x24, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x04},
			ExpectedError: "[]byte length not long enough to contain data length specified in packet",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtCommand{}
			err := got.UnmarshalBinary(test.Data)
			require.Error(t, err)
			require.EqualError(t, err, test.ExpectedError)
		})
	}
}

func TestGoodArtCommandMarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Expected []byte
		Data     *artnet.ArtCommand
	}{
		{
			Name:     "zero",
			Expected: []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0x24, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00},
			Data:     &artnet.ArtCommand{},
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

func TestBadArtCommandMarshal(t *testing.T) {
	tests := []struct {
		Name          string
		Data          *artnet.ArtCommand
		ExpectedError string
	}{
		{
			Name: "data too long",
			Data: &artnet.ArtCommand{
				Data: string(make([]uint8, 513)),
			},
			ExpectedError: "data length must be less than or equal to 512 bytes",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := test.Data.MarshalBinary()
			require.Error(t, err)
			require.EqualError(t, err, test.ExpectedError)
		})
	}
}

func BenchmarkArtCommandMarshalBinary(b *testing.B) {
	data := artnet.ArtCommand{}

	for b.Loop() {
		_, err := data.MarshalBinary()
		require.NoError(b, err)
	}
}
