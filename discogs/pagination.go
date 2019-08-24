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
	PerPage   int    `json:"per_page,omitempty"`
	Page      int    `json:"page,omitempty"`
	Sort      string `json:"sort,omitonempty"`
	SortOrder string `json:"sort_order,omitonempty"`
}
