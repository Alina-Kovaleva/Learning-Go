package main

import (
	"tsi.co/go-chi-sakila/database"
	"tsi.co/go-chi-sakila/resources/models"
	"tsi.co/go-chi-sakila/server"
)

func main() {
	database.Init()
	database.DB.AutoMigrate(&models.Actor{}, &models.Film{}, &models.FilmActor{})

	server.Init()
}
