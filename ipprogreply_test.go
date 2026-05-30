package artnet_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jwetzell/artnet-go"
)

func TestGoodArtIpProgReplyUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtIpProgReply
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtIpProgReply{}

			err := got.UnmarshalBinary(test.Data)
			if err != nil {
				t.Fatalf("failed to Unmarshal ArtIpProgReply: %s", err)
			}
			diff := cmp.Diff(test.Expected, got, cmpopts.IgnoreUnexported(artnet.ArtIpProgReply{}))
			if diff != "" {
				t.Fatalf("ArtIpProgReply does not match\n%s", diff)
			}
		})
	}
}

func BenchmarkArtIpProgReplyMarshalBinary(b *testing.B) {
	data := artnet.ArtIpProgReply{}

	for b.Loop() {
		_, err := data.MarshalBinary()
		if err != nil {
			b.Fatalf("failed to encode ArtIpProgReply: %s", err)
		}
	}
}
