package main

import (
	"tsi.co/go-chi-sakila/database"
	"tsi.co/go-chi-sakila/resources/actors"
	"tsi.co/go-chi-sakila/server"
)

func main() {
	database.Init()
	database.DB.AutoMigrate(&actors.Actor{})

	server.Init()
}
