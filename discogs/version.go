package discogs

type Version struct {
	Release
}

type VersionsResponse struct {
	Versions   []Version `json:"versions"`
	Pagination `json:"pagination"`
}
