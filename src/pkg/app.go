package app

// Exercise defines the properties of a gopher to be listed
type Exercise struct {
	ID    string `json:"ID"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}

// Repository provides access to the gopher storage
type ExerciseRepository interface {
	// CreateGopher saves a given gopher
	CreateGopher(g *Exercise) error
	// FetchGophers return all gophers saved in storage
	FetchGophers() ([]*Exercise, error)
}
