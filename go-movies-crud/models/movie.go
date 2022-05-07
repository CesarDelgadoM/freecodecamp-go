package models

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

func NewMovie(id string, isbn string, title string, director *Director) *Movie {
	return &Movie{
		ID:       id,
		Isbn:     isbn,
		Title:    title,
		Director: director,
	}
}
