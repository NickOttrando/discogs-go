package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetArtist(t *testing.T) {
	artist, err := testClient.GetArtist(1713695)
	assert.Nil(t, err)
	assert.NotNil(t, artist.ID)
	assert.NotNil(t, artist.ResourceURL)
	assert.NotNil(t, artist.Name)
}

func TestGetArtistReleases(t *testing.T) {
	artistReleases, err := testClient.GetArtistReleases(1713695)
	releases := artistReleases.Releases
	assert.Nil(t, err)
	assert.NotNil(t, releases)
	assert.NotNil(t, releases[0].ID)
	assert.NotNil(t, releases[0].Title)

	//test pagination
	assert.NotNil(t, artistReleases.PerPage)
	assert.NotNil(t, artistReleases.Pages)
	assert.NotNil(t, artistReleases.URLs)
	assert.NotNil(t, artistReleases.Items)
	assert.Equal(t, artistReleases.CurrentPage, 1)
}
