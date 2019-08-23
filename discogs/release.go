package discogs

type CommunityStat struct {
	InCollection int `json:"in_collection"`
	InWantlist   int `json:"in_wantlist"`
}

type ReleaseStat struct {
	Community CommunityStat `json:"community"`
}

type Release struct {
	ID          int         `json:"id"`
	Status      string      `json:"status"`
	Stats       ReleaseStat `json:"stats"`
	Thumb       string      `json:"thumb"`
	Title       string      `json:"title"`
	Format      string      `json:"format"`
	Label       string      `json:"label"`
	Role        string      `json:"role"`
	Year        int         `json:"year"`
	ResourceURL string      `json:"resource_url"`
	Artist      string      `json:"artist"`
	Type        string      `json:"type"`
}

// todo: support pagination metadata
type ReleasesResponse struct {
	Releases []Release `json:"releases"`
}
