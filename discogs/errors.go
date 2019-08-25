package discogs

import (
	"fmt"
)

type Error struct {
	// Returned from Discogs API response
	Message    string
	StatusCode int
}

func (e Error) Error() string {
	return fmt.Sprintf("Discogs API Error - http status: %d, message: %s",
		e.StatusCode, e.Message)
}
