package discogs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

var baseURL = "api.discogs.com/"

type Client struct {
	// A specified http client, defaults to http.DefaultClient.
	httpClient *http.Client
	// A User-Agent unique to the client using this Go wrapper.
	userAgent string
	// Optionally pass Key and Secret of your own or user token.
	key    string
	secret string
	token  string
}

type ClientOptions struct {
	UserAgent  string
	Token      string
	Key        string
	Secret     string
	HTTPClient *http.Client
}

// Instantiate a new client, only UserAgent is required.
func NewClient(options ClientOptions) (client *Client, err error) {
	if len(options.UserAgent) == 0 {
		return nil, errors.New("UserAgent is required.")
	}

	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{}
	}

	return &Client{
		userAgent:  options.UserAgent,
		token:      options.Token,
		key:        options.Key,
		secret:     options.Secret,
		httpClient: options.HTTPClient,
	}, nil
}

func (c *Client) Call(endpoint string, method string, body []byte, out interface{}) error {
	req, err := c.newRequest(endpoint, method, bytes.NewReader(body), out)
	if err != nil {
		return err
	}

	return c.do(req, out)
}

// newRequest is used by Call to generate a http.Request with appropriate headers.
func (c *Client) newRequest(endpoint string, method string, body io.Reader, out interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, "https://"+string(baseURL)+endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

// do makes the formatted request and unmarshals the response or returns an err
func (c *Client) do(req *http.Request, out interface{}) error {
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	//fmt.Println(string(resBody))

	// Successful response
	if res.StatusCode == 200 || res.StatusCode == 201 || res.StatusCode == 204 {
		return json.Unmarshal(resBody, out)
	}

	return err
}

func (c *Client) get(endpoint string, out interface{}) error {
	return c.Call(endpoint, "GET", nil, out)
}

func (c *Client) post(endpoint string, body []byte, out interface{}) error {
	return c.Call(endpoint, "POST", body, out)
}

func (c *Client) put(endpoint string, body []byte, out interface{}) error {
	return c.Call(endpoint, "PUT", body, out)
}

func (c *Client) delete(endpoint string) error {
	return c.Call(endpoint, "DELETE", nil, nil)
}
