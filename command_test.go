package artnet_test

import (
	"reflect"
	"testing"

	"github.com/jwetzell/artnet-go"
)

func TestGoodArtCommandUnmarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Data     []byte
		Expected *artnet.ArtCommand
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := &artnet.ArtCommand{}

			err := got.UnmarshalBinary(test.Data)
			if err != nil {
				t.Fatalf("failed to Unmarshal ArtCommand: %s", err)
			}
			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("ArtCommand does not match got: %+v expected: %+v", got, test.Expected)
			}
		})
	}
}
