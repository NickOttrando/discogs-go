# Discogs Go Client Library

Go Client Library for [Discogs API](https://www.discogs.com/developers/).

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

This project is currently in pre-release stages. If you'd like to contribute, see below.

Contributing
------------
Issues and pull reuqests welcomed.

To contribute, just pull the project and install the only dependency, 

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
```
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

    // fetch artist profile
    artist, err := discogsClient.GetArtist(1713695)
    fmt.Println(artist.Name)

    // fetch artist releases, see how many pages there are
    artistReleases, err := discogsClient.GetArtistReleases(1713695, nil)
    fmt.Println(artistReleases.Releases[0].Title)
    fmt.Println(artistReleases.Pages)

    moreReleases, err := discogsClient.GetArtistReleases(1713695, &discogs.ListOptions{Page: 2, Sort: "year", PerPage: 5})
    fmt.Println(moreReleases.Releases[0].Title)
    fmt.Println(moreReleases.Page)
}
```

License
-------
MIT
