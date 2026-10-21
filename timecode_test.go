package artnet_test

import (
	"testing"

	"github.com/jwetzell/artnet-go"
	"github.com/stretchr/testify/require"
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
				StreamId: 0,
				Frames:   11,
				Seconds:  17,
				Minutes:  3,
				Hours:    0,
				Type:     0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtTimeCode{}
			err := got.UnmarshalBinary(test.Data)

			require.NoError(t, err)

			require.Equal(t, artnet.OpTimeCode, got.GetOpCode())

			require.Equal(t, test.Expected, got)
		})
	}
}

func TestBadArtTimeCodeUnmarshal(t *testing.T) {
	tests := []struct {
		Name          string
		Data          []byte
		ExpectedError string
	}{
		{
			Name:          "Too Short",
			Data:          []byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 151, 0, 14, 0, 0},
			ExpectedError: "ArtTimeCode packet must be at least 14 bytes long",
		},
		{
			Name:          "Invalid Header",
			Data:          []byte{65, 114, 116, 45, 78, 101, 115, 0, 0, 151, 0, 14, 0, 0, 11, 17, 3, 0, 0},
			ExpectedError: "ID does not match Art-Net ID",
		},
		{
			Name:          "Incorrect OpCode",
			Data:          []byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 150, 0, 14, 0, 0, 11, 17, 3, 0, 0},
			ExpectedError: "packet does not have the correct OpCode for an ArtTimeCode packet",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtTimeCode{}
			err := got.UnmarshalBinary(test.Data)
			require.Error(t, err)
			require.EqualError(t, err, test.ExpectedError)
		})
	}
}

func TestGoodArtTimeCodeMarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Expected []byte
		Data     *artnet.ArtTimeCode
	}{
		{
			Name:     "Basic timecode",
			Expected: []byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 151, 0, 14, 0, 0, 11, 17, 3, 0, 0},
			Data: &artnet.ArtTimeCode{
				StreamId: 0,
				Frames:   11,
				Seconds:  17,
				Minutes:  3,
				Hours:    0,
				Type:     0,
			},
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

func BenchmarkArtTimeCodeUnmarshalBinary(b *testing.B) {
	data := []byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 151, 0, 14, 0, 0, 11, 17, 3, 0, 0}

	for b.Loop() {
		got := artnet.ArtTimeCode{}

		err := got.UnmarshalBinary(data)
		require.NoError(b, err)
	}
}

func BenchmarkArtTimeCodeMarshalBinary(b *testing.B) {
	data := artnet.ArtTimeCode{
		StreamId: 0,
		Frames:   11,
		Seconds:  17,
		Minutes:  3,
		Hours:    0,
		Type:     0,
	}

	for b.Loop() {
		_, err := data.MarshalBinary()
		require.NoError(b, err)
	}
}

func FuzzArtTimeCodeUnmarshalBinary(f *testing.F) {
	f.Add([]byte{65, 114, 116, 45, 78, 101, 116, 0, 0, 151, 0, 14, 0, 0, 11, 17, 3, 0, 0})
	f.Fuzz(func(t *testing.T, data []byte) {
		artTimeCode := artnet.ArtTimeCode{}
		artTimeCode.UnmarshalBinary(data)
	})
}
