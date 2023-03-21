package main

import (
	"github.com/brunosantanati/api-go-gin/database"
	"github.com/brunosantanati/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
