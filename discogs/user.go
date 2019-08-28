package discogs

type UserStats struct {
	Rating string  `json:"rating"`
	Total  int     `json:"total"`
	Stars  float32 `json:"stars"`
}

type UserData struct {
	InCollection bool `json:"in_collection"`
	InWantlist   bool `json:"in_wantlist"`
}

type User struct {
	ID          int64     `json:"id"`
	UserName    string    `json:"username"`
	Email       string    `json:"email"`
	Stats       UserStats `json:"stats"`
	UID         int64     `json:"uid"`
	URL         string    `json:"url"`
	HtmlURL     string    `json:"html_url"`
	Payment     string    `json:"payment"`
	ResourceURL string    `json:"resource_url"`
}
