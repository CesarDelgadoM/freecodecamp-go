package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CesarDelgadoM/freecodecamp-go/go-movies-crud/controller"
)

type Server struct {
	controller *controller.Controller
}

func NewServer() *Server {
	return &Server{
		controller: controller.NewController(),
	}
}

func (server *Server) Initialize() {
	server.controller.RoutesMovies()
}

func (server *Server) Run(port string) {
	fmt.Printf("Starting server at port: %s\n", port)
	log.Fatal(http.ListenAndServe((":" + port), server.controller.Router))
}
