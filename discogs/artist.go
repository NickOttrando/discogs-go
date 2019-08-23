package discogs

type ArtistAlias struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

type ArtistImage struct {
	URI         string `json:"uri"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	ResourceURL string `json:"resource_url"`
	Type        string `json:"type"`
	URI150      string `json:"string"`
}

type Artist struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	RealName    string        `json:"realname"`
	Profile     string        `json:"profile"`
	ReleasesURL string        `json:"releases_url"`
	URI         string        `json:"uri"`
	ResourceURL string        `json:"resource_url"`
	URLs        []string      `json:"urls"`
	Images      []ArtistImage `json:"images"`
	Aliases     []ArtistAlias `json:"aliases"`
}

func (c *Client) GetArtist(artistID string) (out *Artist, err error) {
	err = c.get("artists/"+artistID, &out)
	return
}

// get releases for artist, todo: support sort params
func (c *Client) GetArtistReleases(artistID string) (out *ReleasesResponse, err error) {
	err = c.get("artists/"+artistID+"/releases", &out)
	return
}
