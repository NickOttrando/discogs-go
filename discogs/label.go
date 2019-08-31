package discogs

import (
	"fmt"
)

// The Label resource represents a label, company, recording studio, location,
// or other entity involved with Artists and Releases. Labels were recently expanded
// in scope to include things that aren’t labels – the name is an artifact of this history.
type Label struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	EntityType     string `json:"entity_type"`
	CatNo          string `json:"catno"`
	ResourceURL    string `json:"resourceURL"`
	EntityTypeName string `json:"entity_type_name"`
}

// Get a label
func (c *Client) GetLabel(labelID int64) (out *Label, err error) {
	err = c.get(fmt.Sprintf("labels/%d", labelID), nil, &out)
	return
}

// Returns a list of Releases associated with the label.
// Accepts Pagination parameters.
func (c *Client) GetLabelReleases(labelID int64, opts *ListOptions) (out *ReleasesResponse, err error) {
	var fmtedOpts *ListOptionsFmted
	if opts != nil {
		fmtedOpts, err = opts.Format()
	}
	if err != nil {
		return nil, err
	}
	err = c.get(fmt.Sprintf("labels/%d/releases", labelID), &fmtedOpts, &out)
	return
}
