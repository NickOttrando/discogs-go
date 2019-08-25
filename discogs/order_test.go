package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

var testOrderID = "1737877-1"

func TestGetOrderNoAuth(t *testing.T) {
	order, err := testClient.GetOrder(testOrderID)
	assert.NotNil(t, err)
	assert.Nil(t, order)
}
