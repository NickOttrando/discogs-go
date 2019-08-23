package discogs

import (
	"fmt"
)

type CommunityStat struct {
	InCollection int `json:"in_collection"`
	InWantlist   int `json:"in_wantlist"`
}

type ReleaseStat struct {
	Community CommunityStat `json:"community"`
}

type ReleaseVideo struct {
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	Embed       bool   `json:"embed"`
	URI         string `json:"uri"`
	Title       string `json:"title"`
}

type Release struct {
	ID          int64          `json:"id"`
	Status      string         `json:"status"`
	Stats       ReleaseStat    `json:"stats"`
	Thumb       string         `json:"thumb"`
	Title       string         `json:"title"`
	Format      string         `json:"format"`
	Label       string         `json:"label"`
	Role        string         `json:"role"`
	Year        int            `json:"year"`
	ResourceURL string         `json:"resource_url"`
	Artist      string         `json:"artist"`
	Type        string         `json:"type"`
	NumForSale  int            `json:"num_for_sale"`
	Styles      []string       `json:"styles"`
	Videos      []ReleaseVideo `json:"videos"`
	Labels      []Label        `json:"labels"`
	Artists     []Artist       `json:"artists"`
}

type ReleasesResponse struct {
	Releases   []Release `json:"releases"`
	Pagination `json:"pagination"`
}

// todo: support curr_abbr toggle
func (c *Client) GetRelease(releaseID int64) (out *Release, err error) {
	err = c.get(fmt.Sprintf("releases/%d", releaseID), &out)
	return
}
