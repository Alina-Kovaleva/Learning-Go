package server

import (
	"github.com/go-chi/chi/v5"
	"tsi.co/go-chi-sakila/resources/actors"
	"tsi.co/go-chi-sakila/resources/filmactor"
	"tsi.co/go-chi-sakila/resources/films"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Mount("/actors", actors.Routes())
	router.Mount("/films", films.Routes())
	router.Mount("/filmactor", filmactor.Routes())

	return router
}
