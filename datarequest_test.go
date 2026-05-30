package artnet_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jwetzell/artnet-go"
)

func TestGoodArtDataRequestUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtDataRequest
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtDataRequest{}

			err := got.UnmarshalBinary(test.Data)
			if err != nil {
				t.Fatalf("failed to Unmarshal ArtDataRequest: %s", err)
			}
			diff := cmp.Diff(test.Expected, got, cmpopts.IgnoreUnexported(artnet.ArtDataRequest{}))
			if diff != "" {
				t.Fatalf("ArtDataRequest does not match\n%s", diff)
			}
		})
	}
}
