package artnet_test

import (
	"testing"

	"github.com/jwetzell/artnet-go"
	"github.com/stretchr/testify/require"
)

func TestGoodArtIpProgReplyUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtIpProgReply
	}{
		{
			Name:     "zero",
			Data:     []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0xf9, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			Expected: &artnet.ArtIpProgReply{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtIpProgReply{}

			err := got.UnmarshalBinary(test.Data)
			require.NoError(t, err)

			require.Equal(t, artnet.OpIpProgReply, got.GetOpCode())

			require.Equal(t, test.Expected, got)
		})
	}
}

func TestBadArtIpProgReplyUnmarshal(t *testing.T) {
	tests := []struct {
		Name          string
		Data          []byte
		ExpectedError string
	}{
		{
			Name:          "Too Short",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0xf9, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			ExpectedError: "ArtIpProgReply packet must be at least 34 bytes long",
		},
		{
			Name:          "Invalid Header",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x73, 0x00, 0x00, 0xf9, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			ExpectedError: "ID does not match Art-Net ID",
		},
		{
			Name:          "Incorrect OpCode",
			Data:          []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0xf8, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			ExpectedError: "packet does not have the correct OpCode for an ArtIpProgReply packet",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtIpProgReply{}
			err := got.UnmarshalBinary(test.Data)
			require.Error(t, err)
			require.EqualError(t, err, test.ExpectedError)
		})
	}
}

func TestGoodArtIpProgReplyMarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Expected []byte
		Data     *artnet.ArtIpProgReply
	}{
		{
			Name:     "zero",
			Expected: []byte{0x41, 0x72, 0x74, 0x2d, 0x4e, 0x65, 0x74, 0x00, 0x00, 0xf9, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			Data:     &artnet.ArtIpProgReply{},
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

func TestBadArtIpProgReplyMarshal(t *testing.T) {
	tests := []struct {
		Name          string
		Data          *artnet.ArtIpProgReply
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

func BenchmarkArtIpProgReplyMarshalBinary(b *testing.B) {
	data := artnet.ArtIpProgReply{}

	for b.Loop() {
		_, err := data.MarshalBinary()
		require.NoError(b, err)
	}
}
