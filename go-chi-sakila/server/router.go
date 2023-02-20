package server

import (
	"github.com/go-chi/chi/v5"
	"tsi.co/go-chi-sakila/resources/actors"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Mount("/actors", actors.Routes())

	return router
}
