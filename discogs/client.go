package discogs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var baseURL = "api.discogs.com/"

// Only userAgent is required to create a client
type Client struct {
	httpClient *http.Client
	userAgent  string
	token      string
}

type ClientOptions struct {
	UserAgent  string
	Token      string
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
		httpClient: options.HTTPClient,
	}, nil
}

func (c *Client) Call(endpoint string, method string, in, out interface{}) error {
	req, err := c.newRequest(endpoint, method, in, out)
	if err != nil {
		return err
	}

	return c.do(req, out)
}

// newRequest is used by Call to generate a http.Request with appropriate headers.
func (c *Client) newRequest(endpoint string, method string, in, out interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, "https://"+string(baseURL)+endpoint, nil)
	if err != nil {
		return nil, err
	}

	switch in := in.(type) {
	case nil: // nop
	default:
		buf, err := json.Marshal(in)
		if err != nil {
			return nil, err
		}
		m := make(map[string]string)
		err = json.Unmarshal(buf, &m)
		q := req.URL.Query()
		for k, v := range m {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.userAgent)
	if len(c.token) > 0 {
		req.Header.Add("Authorization", "Discogs token:"+c.token)
	}

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

func (c *Client) get(endpoint string, in, out interface{}) error {
	return c.Call(endpoint, "GET", in, out)
}

func (c *Client) post(endpoint string, in, out interface{}) error {
	return c.Call(endpoint, "POST", in, out)
}

func (c *Client) put(endpoint string, in, out interface{}) error {
	return c.Call(endpoint, "PUT", in, out)
}

func (c *Client) delete(endpoint string) error {
	return c.Call(endpoint, "DELETE", nil, nil)
}
