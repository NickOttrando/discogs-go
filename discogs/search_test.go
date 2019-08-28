package discogs

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSearchNoAuth(t *testing.T) {
	results, err := testClient.Search(&SearchParams{Type: "label", Query: "Warp Records"}, nil)
	assert.NotNil(t, err)
	assert.Nil(t, results)
}
