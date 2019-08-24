package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

var testLabelID = int64(209402)

func TestGetLabel(t *testing.T) {
	label, err := testClient.GetLabel(testLabelID)
	assert.Nil(t, err)
	assert.NotNil(t, label.ID)
	assert.Equal(t, label.ID, testLabelID)
	assert.NotNil(t, label.Name)
}

func TestGetLabelReleases(t *testing.T) {
	labelReleases, err := testClient.GetLabelReleases(testLabelID, nil)
	releases := labelReleases.Releases
	assert.Nil(t, err)
	assert.NotNil(t, releases)
	assert.NotNil(t, releases[0].ID)
	assert.NotNil(t, releases[0].Title)

	//test pagination
	assert.NotNil(t, labelReleases.PerPage)
	assert.NotNil(t, labelReleases.Pages)
	assert.NotNil(t, labelReleases.URLs)
	assert.NotNil(t, labelReleases.Items)
	assert.Equal(t, labelReleases.Page, 1)

    nextReleases, err := testClient.GetLabelReleases(testLabelID, &PageOptions{Page: "2"})
    assert.NotNil(t, nextReleases.Releases)
    assert.Equal(t, nextReleases.Page, 2)
}
