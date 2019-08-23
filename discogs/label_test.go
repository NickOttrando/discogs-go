package discogs

import (
    "testing"

    assert "github.com/stretchr/testify/require"
)

func TestGetLabel(t *testing.T) {
    label, err := testClient.GetLabel(209402)
    assert.Nil(t, err)
    assert.NotNil(t, label.ID)
    assert.Equal(t, label.ID, int64(209402))
    assert.NotNil(t, label.Name)
}

func TestGetLabelReleases(t *testing.T) {
    labelReleases, err := testClient.GetLabelReleases(209402)
    releases := labelReleases.Releases
    assert.Nil(t, err)
    assert.NotNil(t, releases)
    assert.NotNil(t, releases[0].ID)
    assert.NotNil(t, releases[0].Title)
}
