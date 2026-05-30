package artnet_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jwetzell/artnet-go"
)

func TestGoodArtTriggerUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtTrigger
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtTrigger{}

			err := got.UnmarshalBinary(test.Data)
			if err != nil {
				t.Fatalf("failed to Unmarshal ArtTrigger: %s", err)
			}
			diff := cmp.Diff(test.Expected, got, cmpopts.IgnoreUnexported(artnet.ArtTrigger{}))
			if diff != "" {
				t.Fatalf("ArtTrigger does not match\n%s", diff)
			}
		})
	}
}

func BenchmarkArtTriggerMarshalBinary(b *testing.B) {
	data := artnet.ArtTrigger{}

	for b.Loop() {
		_, err := data.MarshalBinary()
		if err != nil {
			b.Fatalf("failed to encode ArtTrigger: %s", err)
		}
	}
}
