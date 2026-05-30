package artnet_test

import (
	"reflect"
	"testing"

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
			if !reflect.DeepEqual(got, test.Expected) {
				t.Fatalf("ArtPollReply does not match got: %+v expected: %+v", got, test.Expected)
			}
		})
	}
}
