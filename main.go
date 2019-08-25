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
