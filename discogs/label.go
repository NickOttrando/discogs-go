package discogs

import (
	"fmt"
)

type Label struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	EntityType     string `json:"entity_type"`
	CatNo          string `json:"catno"`
	ResourceURL    string `json:"resourceURL"`
	EntityTypeName string `json:"entity_type_name"`
}

func (c *Client) GetLabel(labelID int64) (out *Label, err error) {
	err = c.get(fmt.Sprintf("labels/%d", labelID), &out)
	return
}

func (c *Client) GetLabelReleases(labelID int64) (out *ReleasesResponse, err error) {
	err = c.get(fmt.Sprintf("labels/%d/releases", labelID), &out)
	return
}
