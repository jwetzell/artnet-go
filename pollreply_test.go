package artnet_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jwetzell/artnet-go"
)

func TestGoodArtPollReplyUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtPollReply
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtPollReply{}

			err := got.UnmarshalBinary(test.Data)
			if err != nil {
				t.Fatalf("failed to Unmarshal ArtPollReply: %s", err)
			}
			diff := cmp.Diff(test.Expected, got, cmpopts.IgnoreUnexported(artnet.ArtPollReply{}))
			if diff != "" {
				t.Fatalf("ArtPollReply does not match\n%s", diff)
			}
		})
	}
}
