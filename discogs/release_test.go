package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetRelease(t *testing.T) {
	release, err := testClient.GetRelease(10201544)
	assert.Nil(t, err)
	assert.NotNil(t, release.ID)
	assert.Equal(t, release.ID, int64(10201544))
	assert.NotNil(t, release.Title)
}
