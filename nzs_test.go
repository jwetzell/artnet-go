package artnet_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jwetzell/artnet-go"
)

func TestGoodArtNzsUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtNzs
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtNzs{}

			err := got.UnmarshalBinary(test.Data)
			if err != nil {
				t.Fatalf("failed to Unmarshal ArtNzs: %s", err)
			}
			diff := cmp.Diff(test.Expected, got, cmpopts.IgnoreUnexported(artnet.ArtNzs{}))
			if diff != "" {
				t.Fatalf("ArtNzs does not match\n%s", diff)
			}
		})
	}
}
