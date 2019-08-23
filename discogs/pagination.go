package discogs

type PaginationURL struct {
	Last string `json:"last"`
	Next string `json:"next"`
}

type Pagination struct {
	PerPage     int           `json:"per_page"`
	Pages       int           `json:"pages"`
	CurrentPage int           `json:"page"`
	URLs        PaginationURL `json:"urls"`
	Items       int           `json:"items"`
}
