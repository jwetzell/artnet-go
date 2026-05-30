package artnet_test

import (
	"reflect"
	"testing"

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
			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("ArtNzs does not match got: %+v expected: %+v", got, test.Expected)
			}
		})
	}
}
