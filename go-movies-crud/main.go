package main

import "github.com/CesarDelgadoM/freecodecamp-go/go-movies-crud/server"

func main() {
	server := server.NewServer()
	server.Initialize()
	server.Run("8081")
}
