package discogs

import (
	"fmt"
)

// The Master resource represents a set of similar Releases.
// Masters (also known as “master releases”) have a “main release” which is
// often the chronologically earliest.
type Master struct {
	Release
}

// Get a master release
func (c *Client) GetMaster(masterID int64) (out *Master, err error) {
	err = c.get(fmt.Sprintf("masters/%d", masterID), nil, &out)
	return
}

// Retrieves a list of all Releases that are versions of this master.
// Accepts Pagination parameters.
func (c *Client) GetMasterVersions(masterID int64, opts *ListOptions) (out *VersionsResponse, err error) {
	var fmtedOpts *ListOptionsFmted
	if opts != nil {
		fmtedOpts, err = opts.Format()
	}
	if err != nil {
		return nil, err
	}
	err = c.get(fmt.Sprintf("masters/%d/versions", masterID), &fmtedOpts, &out)
	return
}
