package main

import (
	"github.com/maksimfisenko/moxer/internal/server"

	_ "github.com/maksimfisenko/moxer/docs"
)

//	@title			Moxer API
//	@version		1.0
//	@description	This is a backend API of the Moxer application.

//	@contact.name	Maksim Fisenko
//	@contact.email	fisenkomaksim.id@gmail.com

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@accept		json
//	@produce	json

func main() {
	server.Start()
}
