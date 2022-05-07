package service

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/CesarDelgadoM/freecodecamp-go/go-movies-crud/models"
	"github.com/CesarDelgadoM/freecodecamp-go/go-movies-crud/repository"
	"github.com/CesarDelgadoM/freecodecamp-go/go-movies-crud/response"
	"github.com/gorilla/mux"
)

type Service struct {
	repo *repository.Repository
}

func NewService() *Service {
	return &Service{
		repo: repository.NewRepository(),
	}
}

func (service *Service) GetAllMovies(rw http.ResponseWriter, r *http.Request) {
	movies := service.repo.Movies
	response.Json(rw, http.StatusOK, movies)
}

func (service *Service) GetMovie(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, movie := range service.repo.Movies {
		if movie.ID == id {
			response.Json(rw, http.StatusFound, movie)
			return
		}
	}
}

func (service *Service) SaveMovie(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	var movie models.Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		response.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	service.repo.Movies = append(service.repo.Movies, movie)
	response.Json(rw, http.StatusCreated, movie)
}
