package discogs

type Artist struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Profile string `json:"profile"`
}

func (c *Client) GetArtist(artistID string) (out *Artist, err error) {
	err = c.get("artists/"+artistID, &out)
	return
}
