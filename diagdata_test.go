package artnet_test

import (
	"reflect"
	"testing"

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
			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("ArtDiagData does not match got: %+v expected: %+v", got, test.Expected)
			}
		})
	}
}
