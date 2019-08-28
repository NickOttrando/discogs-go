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
