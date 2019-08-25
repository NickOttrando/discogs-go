package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

var testReleaseID = int64(28060)
var testUsername = "pusscadour"

func TestGetRelease(t *testing.T) {
	release, err := testClient.GetRelease(testReleaseID)
	assert.Nil(t, err)
	assert.NotNil(t, release.ID)
	assert.Equal(t, release.ID, testReleaseID)
	assert.NotNil(t, release.Title)
}

func TestGetReleaseRating(t *testing.T) {
	releaseRating, err := testClient.GetReleaseRating(testReleaseID)
	rating := releaseRating.Rating
	assert.Nil(t, err)
	assert.NotNil(t, releaseRating)
	assert.Equal(t, releaseRating.ReleaseID, testReleaseID)
	assert.NotNil(t, rating)
	assert.NotNil(t, rating.Count)
	assert.NotNil(t, rating.Average)
}

func TestGetReleaseRatingByUser(t *testing.T) {
	userReleaseRating, err := testClient.GetReleaseRatingByUser(testReleaseID, testUsername)
	assert.Nil(t, err)
	assert.NotNil(t, userReleaseRating)
	assert.NotNil(t, userReleaseRating.Count)
	assert.NotNil(t, userReleaseRating.Average)
}
