package discogs

type PaginationURL struct {
	Last string `json:"last"`
	Next string `json:"next"`
}

type Pagination struct {
	PerPage int           `json:"per_page"`
	Pages   int           `json:"pages"`
	Page    int           `json:"page"`
	URLs    PaginationURL `json:"urls"`
	Items   int           `json:"items"`
}

type PageOptions struct {
	PerPage   string `json:"per_page,omitempty"`
	Page      string `json:"page,omitempty"`
	Sort      string `json:"sort,omitempty"`
	SortOrder string `json:"sort_order,omitempty"`
}
