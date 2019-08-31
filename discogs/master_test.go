package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

var testMasterID = int64(565)

func TestGetMaster(t *testing.T) {
	master, err := testClient.GetMaster(testMasterID)
	assert.Nil(t, err)
	assert.NotNil(t, master.ID)
}

func TestGetMasterVersions(t *testing.T) {
	masterVersions, err := testClient.GetMasterVersions(testMasterID, nil)
	versions := masterVersions.Versions
	assert.Nil(t, err)
	assert.NotNil(t, versions)
	assert.NotNil(t, versions[0].Title)
}
