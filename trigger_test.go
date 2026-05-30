package artnet_test

import (
	"reflect"
	"testing"

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
			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("ArtTrigger does not match got: %+v expected: %+v", got, test.Expected)
			}
		})
	}
}
