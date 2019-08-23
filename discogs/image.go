package discogs

type Image struct {
	URI         string `json:"uri"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	ResourceURL string `json:"resource_url"`
	Type        string `json:"type"`
	URI150      string `json:"string"`
}
