package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetMaster(t *testing.T) {
	master, err := testClient.GetMaster(565)
	assert.Nil(t, err)
	assert.NotNil(t, master.ID)
}

func TestGetMasterVersions(t *testing.T) {
	masterVersions, err := testClient.GetMasterVersions(565)
	versions := masterVersions.Versions
	assert.Nil(t, err)
	assert.NotNil(t, versions)
	assert.NotNil(t, versions[0].Title)
}
