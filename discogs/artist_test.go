package discogs

import (
	"testing"
	//"fmt"

	assert "github.com/stretchr/testify/require"
)

var testArtistID = int64(1713695)

func TestGetArtist(t *testing.T) {
	artist, err := testClient.GetArtist(testArtistID)
	assert.Nil(t, err)
	assert.NotNil(t, artist.ID)
	assert.NotNil(t, artist.ResourceURL)
	assert.NotNil(t, artist.Name)
}

func TestGetArtistReleases(t *testing.T) {
	artistReleases, err := testClient.GetArtistReleases(testArtistID, nil)
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
	assert.Equal(t, artistReleases.Page, 1)

	nextReleases, err := testClient.GetArtistReleases(testArtistID, &ListOptions{Page: 2, Sort: "year", PerPage: 5})
	assert.NotNil(t, nextReleases.Releases)
	assert.Equal(t, nextReleases.Page, 2)
	assert.Len(t, nextReleases.Releases, 5)
}
