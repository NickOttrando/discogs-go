package discogs

import (
	"testing"
    "fmt"

	assert "github.com/stretchr/testify/require"
)

func TestGetArtist(t *testing.T) {
    artist, err := testClient.GetArtist("1713695")
    fmt.Println(artist.Name)
    assert.Nil(t, err)
	assert.NotNil(t, artist)
}
