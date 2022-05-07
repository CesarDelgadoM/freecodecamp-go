package controller

import (
	"github.com/CesarDelgadoM/freecodecamp-go/go-movies-crud/middleware"
	"github.com/CesarDelgadoM/freecodecamp-go/go-movies-crud/service"
	"github.com/gorilla/mux"
)

type Controller struct {
	service *service.Service
	Router  *mux.Router
}

func NewController() *Controller {
	return &Controller{
		service: service.NewService(),
		Router:  mux.NewRouter(),
	}
}

func (controller *Controller) RoutesMovies() {
	controller.Router.HandleFunc("/movie/getall",
		middleware.SetMiddlewareJson(controller.service.GetAllMovies)).Methods("GET")
	controller.Router.HandleFunc("/movie/get/{id}",
		middleware.SetMiddlewareJson(controller.service.GetMovie)).Methods("GET")
	controller.Router.HandleFunc("/movie/save",
		middleware.SetMiddlewareJson(controller.service.SaveMovie)).Methods("POST")
}
