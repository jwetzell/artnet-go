package artnet_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jwetzell/artnet-go"
)

func TestGoodArtDiagDataUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtDiagData
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtDiagData{}

			err := got.UnmarshalBinary(test.Data)
			if err != nil {
				t.Fatalf("failed to Unmarshal ArtDiagData: %s", err)
			}
			diff := cmp.Diff(test.Expected, got, cmpopts.IgnoreUnexported(artnet.ArtDiagData{}))
			if diff != "" {
				t.Fatalf("ArtDiagData does not match\n%s", diff)
			}
		})
	}
}

func BenchmarkArtDiagDataMarshalBinary(b *testing.B) {
	data := artnet.ArtDiagData{}

	for b.Loop() {
		_, err := data.MarshalBinary()
		if err != nil {
			b.Fatalf("failed to encode ArtDiagData: %s", err)
		}
	}
}
