package discogs

import (
	"fmt"
)

type SearchParams struct {
	Query        string `json:"q,omitempty"`
	Type         string `json:"type,omitempty"`
	Title        string `json:"title,omitempty"`
	ReleaseTitle string `json:"release_title,omitempty"`
	Credit       string `json:"credit,omitempty"`
	Artist       string `json:"artist,omitempty"`
	Anv          string `json:"anv,omitempty"`
	Label        string `json:"label,omitempty"`
	Genre        string `json:"genre,omitempty"`
	Style        string `json:"style,omitempty"`
	Country      string `json:"country,omitempty"`
	Year         string `json:"year,omitempty"`
	Format       string `json:"format,omitempty"`
	CatNo        string `json:"catno,omitempty"`
	Barcode      string `json:"barcode,omitempty"`
	Track        string `json:"track,omitempty"`
	Submitter    string `json:"submitter,omitempty"`
	Contributor  string `json:"contributor,omitempty"`
}

type SearchResult struct {
	ID          int64    `json:"id"`
	Thumb       string   `json:"thumb"`
	Title       string   `json:"title"`
	UserData    UserData `json:"user_data"`
	MasterURL   string   `json:"master_url"`
	URI         string   `json:"uri"`
	CoverImage  string   `json:"cover_image"`
	ResourceURL string   `json:"resource_url"`
	Type        string   `json:"type"`
}

type SearchResponse struct {
	Results []SearchResult `json:"results"`
	Pagination
}

// Issue a search query to the Discogs database.
// This endpoint accepts pagination parameters.
// Authentication (as any user) is required.
func (c *Client) Search(query *SearchParams, opts *ListOptions) (out *SearchResponse, err error) {
	var fmtedOpts *ListOptionsFmted
	if opts != nil {
		fmtedOpts, err = opts.Format()
	}
	if err != nil {
		return nil, err
	}

	err = c.searchGet(fmt.Sprintf("database/search"), &fmtedOpts, &query, &out)
	return
}
