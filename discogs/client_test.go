package discogs

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

var (
	userAgent = "Discogs Go Test Client"
	token     = "getenv tok"
	key       = "getenv key"
	secret    = "getenv secret"
)

var testOptions = ClientOptions{
	userAgent,
	token,
	key,
	secret,
	&http.Client{},
}

var testClient, err = NewClient(testOptions)

func TestClient(t *testing.T) {
	assert, require := assert.New(t), require.New(t)

	require.NoError(err)
	assert.Equal(testClient.userAgent, testOptions.UserAgent)
}
