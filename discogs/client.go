package discogs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var baseURL = "api.discogs.com/"

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

// Builds request object, calls do once it's been formatted
func (c *Client) Call(endpoint string, method string, listOpts, queryArgs, out interface{}) error {

	req, err := http.NewRequest(method, "https://"+string(baseURL)+endpoint, nil)
	if err != nil {
		return err
	}

	// Query args strings for pagination, sort, etc.
	if listOpts != nil {
		addQueryArgs(req, listOpts)
	}

	if queryArgs != nil {
		addQueryArgs(req, queryArgs)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.userAgent)
	if len(c.token) > 0 {
		req.Header.Add("Authorization", "Discogs token="+c.token)
	}
	return c.do(req, out)
}

func addQueryArgs(req *http.Request, args interface{}) error {
	buf, err := json.Marshal(args)
	if err != nil {
		return err
	}
	m := make(map[string]string)

	err = json.Unmarshal(buf, &m)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	for k, v := range m {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return nil
}

// Makes the formatted request and handles the response
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

	// If we need to check responses
	//fmt.Println(req.URL.String())
	//fmt.Println(string(resBody))

	// Successful response
	if res.StatusCode == 200 || res.StatusCode == 201 || res.StatusCode == 204 {
		return json.Unmarshal(resBody, out)
	}

	// Attempt to unmarshal into error
	var discogsErr Error
	if err = json.Unmarshal(resBody, &discogsErr); err != nil {
		return err
	}
	discogsErr.StatusCode = res.StatusCode
	return discogsErr
}

func (c *Client) get(endpoint string, listOpts, out interface{}) error {
	return c.Call(endpoint, "GET", listOpts, nil, out)
}

func (c *Client) searchGet(endpoint string, listOpts, queryArgs, out interface{}) error {
	return c.Call(endpoint, "GET", listOpts, queryArgs, out)
}
