package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`                 // Time when the movie was added to the database
	Title     string    `json:"Title"`             // Movie title
	Year      int32     `json:"year,omitempty"`    // Movie release year
	Runtime   Runtime   `json:"runtime,omitempty"` // Movie runtime (in minutes)
	Genres    []string  `json:"genres,omitempty"`  // Slice of genres (action, comedy, ...)
	Version   int32     `json:"version"`           // Starts at 1, incremented each time the information is changed
}
