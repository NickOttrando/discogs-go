package discogs

import (
	"fmt"
)

type ArtistAlias struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

type Artist struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	RealName    string        `json:"realname"`
	Profile     string        `json:"profile"`
	ReleasesURL string        `json:"releases_url"`
	URI         string        `json:"uri"`
	ResourceURL string        `json:"resource_url"`
	URLs        []string      `json:"urls"`
	Images      []Image       `json:"images"`
	Aliases     []ArtistAlias `json:"aliases"`
}

func (c *Client) GetArtist(artistID int64) (out *Artist, err error) {
	err = c.get(fmt.Sprintf("artists/%d", artistID), &out)
	return
}

// get releases for artist, todo: support sort params
func (c *Client) GetArtistReleases(artistID int64) (out *ReleasesResponse, err error) {
	err = c.get(fmt.Sprintf("artists/%d/releases", artistID), &out)
	return
}
