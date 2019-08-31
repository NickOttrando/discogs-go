# Discogs Go Client Library
[![Build Status](https://travis-ci.org/NickOttrando/discogs-go.svg?branch=master)](https://travis-ci.org/NickOttrando/discogs-go)

Go Client Library for [Discogs REST API](https://www.discogs.com/developers/).

Discogs is a massive database for music releases, and a marketplace to buy and sell physical music.

This wrapper currently supports the following (* requires auth through Token)
   - Artists
   - Releases
   - Master Releases
   - Labels
   - Inventory
   - Listings
   - Orders*
   
   - Pagination / Sorting

Contributing
------------
Issues and pull reuqests welcomed.

To contribute, just pull the project and install the only [dependencies](https://github.com/stretchr/testify), testify/assert and testify/require. 

Testing
-------
`make test`
   
Installation
------------
`go get -u github.com/NickOttrando/discogs-go`

Documentation
-------------
See Discogs API docs [here](https://www.discogs.com/developers/).

Getting Started
---------------
Below is an example of how to get started and how to use pagination/sorting for list endpoints. For more examples, see the test files. Only a UserAgent is required unless hitting Auth endpoints (orders, etc.)
```go
package main

import (
	"fmt"
	"os"

	"github.com/NickOttrando/discogs-go/discogs"
)

func main() {
	clientOptions := discogs.ClientOptions{
		UserAgent: os.Getenv("DISCOGS_USER_AGENT"),
		Token:     os.Getenv("DISCOGS_TOKEN"),
	}

	discogsClient, err := discogs.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	// search for a label
	labelSearch, err := discogsClient.Search(&discogs.SearchParams{Type: "label", Query: "Warp Records"}, nil)
	fmt.Println(labelSearch.Results[0].Title)

	// search for an artist
	artistSearch, err := discogsClient.Search(&discogs.SearchParams{Type: "artist", Query: "Aphex Twin"}, &discogs.ListOptions{PerPage: 1})
	artistID := artistSearch.Results[0].ID
	fmt.Println(artistID)

	// fetch artist profile
	artist, err := discogsClient.GetArtist(artistID)
	fmt.Println(artist.Name)

	// fetch artist releases, see how many pages there are
	artistReleases, err := discogsClient.GetArtistReleases(artistID, nil)
	fmt.Println(artistReleases.Releases[0].Title)
	fmt.Println(artistReleases.Pages)

	moreReleases, err := discogsClient.GetArtistReleases(artistID, &discogs.ListOptions{Page: 2, Sort: "year", PerPage: 5})
	fmt.Println(moreReleases.Releases[0].Title)
	fmt.Println(moreReleases.Page)

	// fetch label and their releases
	label, err := discogsClient.GetLabel(23528)
	fmt.Println(label.Name)

	labelReleases, err := discogsClient.GetLabelReleases(23528, nil)
	fmt.Println(labelReleases.Releases[0].Title)
}
```

License
-------
MIT
