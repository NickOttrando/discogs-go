package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

var testListingID = int64(988369143)
var testUser = "clintonstreetrecords"

func TestGetListing(t *testing.T) {
	listing, err := testClient.GetListing(testListingID)
	assert.Nil(t, err)
	assert.NotNil(t, listing.ID)
	assert.NotNil(t, listing.Status)
	assert.NotNil(t, listing.AllowOffers)
	assert.NotNil(t, listing.URI)
	assert.NotNil(t, listing.Comments)
	assert.NotNil(t, listing.Seller)
	assert.NotNil(t, listing.Price)
	assert.NotNil(t, listing.OriginalShippingPrice)
	assert.NotNil(t, listing.SleeveCondition)
	assert.NotNil(t, listing.ShippingPrice)
	assert.NotNil(t, listing.Release)
	assert.NotNil(t, listing.ResourceURL)
	assert.NotNil(t, listing.Audio)
	assert.NotNil(t, listing.Condition)
	assert.NotNil(t, listing.Posted)
}

func TestGetInventory(t *testing.T) {
	inventory, err := testClient.GetUserInventory(testUser, &PageOptions{PerPage: "10"})
	listings := inventory.Listings
	assert.Nil(t, err)
	assert.NotNil(t, inventory)
	assert.NotNil(t, listings)
	assert.NotNil(t, inventory.Page)
}
