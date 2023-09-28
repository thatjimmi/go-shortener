package main

import (
	"lenk/model"
	"lenk/server"
)

func main() {
	model.Setup()
	server.SetupRoutes()
}
