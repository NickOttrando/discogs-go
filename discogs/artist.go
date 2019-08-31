package discogs

import (
	"fmt"
)

type ArtistAlias struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

// The Artist resource represents a person in the Discogs database who
// contributed to a Release in some capacity.
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

// Get an artist
func (c *Client) GetArtist(artistID int64) (out *Artist, err error) {
	err = c.get(fmt.Sprintf("artists/%d", artistID), nil, &out)
	return
}

// Get an artist’s Releases and Masters
// Accepts Pagination.
func (c *Client) GetArtistReleases(artistID int64, opts *ListOptions) (out *ReleasesResponse, err error) {
	var fmtedOpts *ListOptionsFmted
	if opts != nil {
		fmtedOpts, err = opts.Format()
	}
	if err != nil {
		return nil, err
	}
	err = c.get(fmt.Sprintf("artists/%d/releases", artistID), &fmtedOpts, &out)
	return
}
