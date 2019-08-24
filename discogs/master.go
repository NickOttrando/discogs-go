package discogs

import (
	"fmt"
)

type Master struct {
	Release
}

func (c *Client) GetMaster(masterID int64) (out *Master, err error) {
	err = c.get(fmt.Sprintf("masters/%d", masterID), nil, &out)
	return
}

func (c *Client) GetMasterVersions(masterID int64) (out *VersionsResponse, err error) {
	err = c.get(fmt.Sprintf("masters/%d/versions", masterID), nil, &out)
	return
}
