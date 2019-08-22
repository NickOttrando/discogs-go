package discogs

import (
    "net/url"
)

var baseURL = &url.URL{
    Scheme: "https",
    Host: "api.discogs.com",
}

type Client struct {
    // A specified http client, defaults to http.DefaultClient.
    httpClient    *http.Client
    // A User-Agent unique to the client using this Go wrapper.
    userAgent     string
    // Optionally pass Key and Secret of your own or user token.
    key           string
    secret        string
    token         string
}

type ClientOptions struct {
	UserAgent   string
    Token       string
	Key         string
	Secret      string
	HTTPClient  *http.Client
}

// Instantiate a new client, only UserAgent is required.
func New(options ClientOptions) (client *Client, err error) {
    if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{}
	}

    return &Client{
        userAgent: options.UserAgent,
        token: options.Token,
        key: options.Key,
        secret: options.Secret,
        httpClient: options.HTTPClient,
    }, nil
}


