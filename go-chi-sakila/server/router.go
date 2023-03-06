package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"tsi.co/go-chi-sakila/resources/actors"
	"tsi.co/go-chi-sakila/resources/filmactor"
	"tsi.co/go-chi-sakila/resources/films"
)

func Router() chi.Router {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"https://*", "http://*"}, AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}, AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, ExposedHeaders: []string{"Link"}, AllowCredentials: false, MaxAge: 300}))
	router.Mount("/actors", actors.Routes())
	router.Mount("/films", films.Routes())
	router.Mount("/filmactor", filmactor.Routes())

	return router
}
