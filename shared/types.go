package shared

// Suggestion represents a simple course suggestion.
type Suggestion struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Position represents a geo coordinate.
type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Route represents part of a course with a position.
type Route struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Position    Position `json:"position"`
}

// Detail represents detailed course information.
type Detail struct {
	Summary string  `json:"summary"`
	Routes  []Route `json:"routes"`
}
