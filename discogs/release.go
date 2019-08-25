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

type Track struct {
	Duration string `json:"duration"`
	Position string `json:"position"`
	Type     string `json:"type_"`
	Title    string `json:"title"`
}

type Rating struct {
	Count   int     `json:"count"`
	Average float32 `json:"average"`
}

type RatingResponse struct {
	Rating    Rating `json:"rating"`
	ReleaseID int64  `json:"release_id"`
}

type Release struct {
	ID                   int64       `json:"id"`
	Status               string      `json:"status"`
	Stats                ReleaseStat `json:"stats"`
	Thumb                string      `json:"thumb"`
	Title                string      `json:"title"`
	Format               string      `json:"format"`
	Label                string      `json:"label"`
	Role                 string      `json:"role"`
	Year                 int         `json:"year"`
	ResourceURL          string      `json:"resource_url"`
	Artist               string      `json:"artist"`
	Type                 string      `json:"type"`
	NumForSale           int         `json:"num_for_sale"`
	DataQuality          string      `json:"data_quality"`
	LowestPrice          float32     `json:"lowest_price"`
	MostRecentReleaseURL string      `json:"most_recent_release_url"`
	MostRecentRelease    int64       `json:"most_recent_release"`
	VersionsURL          string      `json:"versions_url"`
	MainReleaseURL       string      `json:"main_release_url"`
	MainRelease          int64       `json:"main_release"`
	Notes                string      `json:"notes"`
	URI                  string      `json:"uri"`
	Styles               []string    `json:"styles"`
	Images               []Image     `json:"images"`
	Genres               []string    `json:"genres"`
	Videos               []Video     `json:"videos"`
	Labels               []Label     `json:"labels"`
	Artists              []Artist    `json:"artists"`
	Tracklist            []Track     `json:"tracklist"`
}

type ReleasesResponse struct {
	Releases   []Release `json:"releases"`
	Pagination `json:"pagination"`
}

// todo: support curr_abbr toggle
func (c *Client) GetRelease(releaseID int64) (out *Release, err error) {
	err = c.get(fmt.Sprintf("releases/%d", releaseID), nil, &out)
	return
}

func (c *Client) GetReleaseRating(releaseID int64) (out *RatingResponse, err error) {
	err = c.get(fmt.Sprintf("releases/%d/rating", releaseID), nil, &out)
	return
}

func (c *Client) GetReleaseRatingByUser(releaseID int64, username string) (out *Rating, err error) {
	err = c.get(fmt.Sprintf("releases/%d/rating/%s", releaseID, username), nil, &out)
	return
}
