package discogs

type Label struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	EntityType     string `json:"entity_type"`
	CatNo          string `json:"catno"`
	ResourceURL    string `json:"resourceURL"`
	EntityTypeName string `json:"entity_type_name"`
}
