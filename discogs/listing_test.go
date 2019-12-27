package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

var testListingID = int64(988369143)
var testUser = "clintonstreetrecords"

func TestGetListing(t *testing.T) {
	listing, err := testClient.GetListing(testListingID)
	assert.NotNil(t, err)
  assert.Nil(t, listing)
}

func TestGetInventory(t *testing.T) {
	inventory, err := testClient.GetUserInventory(testUser, &ListOptions{PerPage: 10})
	listings := inventory.Listings
	assert.Nil(t, err)
	assert.NotNil(t, inventory)
	assert.NotNil(t, listings)
	assert.NotNil(t, inventory.Page)
}
