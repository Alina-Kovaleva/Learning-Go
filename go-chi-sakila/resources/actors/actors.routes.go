package actors

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListActors)
	router.Post("/", CreateActor)

	router.Get("/{id}", GetActorById)
	router.Delete("/{id}", DeleteActor)
	router.Patch("/{id}", UpdateActor)

	return router
}
