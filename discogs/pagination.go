package discogs

import (
    "strconv"
)

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


type ListOptions struct {
	PerPage   int    `json:"per_page,omitempty"`
	Page      int    `json:"page,omitempty"`
	Sort      string `json:"sort,omitempty"`
	SortOrder string `json:"sort_order,omitempty"`
}

type ListOptionsFmted struct {
	PerPage   string `json:"per_page,omitempty"`
	Page      string `json:"page,omitempty"`
	Sort      string `json:"sort,omitempty"`
	SortOrder string `json:"sort_order,omitempty"`
}

func (o *ListOptions) Format() (out *ListOptionsFmted, err error) {
    //o is nil if no options given, avoid null pointer exception
    if o == nil {
        return
    }
    return &ListOptionsFmted{
        PerPage: strconv.Itoa(o.PerPage),
        Page: strconv.Itoa(o.Page),
        Sort: o.Sort,
        SortOrder: o.SortOrder,
    }, nil
}
