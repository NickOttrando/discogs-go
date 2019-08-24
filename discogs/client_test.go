package discogs

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	userAgent = "Discogs Go Test Client"
	token     = os.Getenv("DISCOGS_TOKEN")
)

var testOptions = ClientOptions{
	userAgent,
	token,
	&http.Client{},
}

var testClient, err = NewClient(testOptions)

func TestClient(t *testing.T) {
	assert, require := assert.New(t), require.New(t)

	require.NoError(err)
	assert.Equal(testClient.userAgent, testOptions.UserAgent)
}
