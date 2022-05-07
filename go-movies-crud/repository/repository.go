package repository

import "github.com/CesarDelgadoM/freecodecamp-go/go-movies-crud/models"

type Repository struct {
	Movies []models.Movie
}

func NewRepository() *Repository {
	return &Repository{
		Movies: loadData(),
	}
}

func loadData() []models.Movie {
	var movies []models.Movie
	movies = append(movies,
		*models.NewMovie("1", "452378", "Movie One", models.NewDirector("Cesar", "Delgado")))
	movies = append(movies,
		*models.NewMovie("2", "895633", "Movie Two", models.NewDirector("Paola", "Avella")))
	movies = append(movies,
		*models.NewMovie("3", "127889", "Movie Three", models.NewDirector("Juan y Jose", "Los perritos")))
	return movies
}
