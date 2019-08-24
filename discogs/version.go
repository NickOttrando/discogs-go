package discogs

//todo: support filter facets and filters?
type Version struct {
	Release
}

type VersionsResponse struct {
	Versions   []Version `json:"versions"`
	Pagination `json:"pagination"`
}
