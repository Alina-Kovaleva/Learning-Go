package filmactor

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", CreateFilmActor)
	router.Delete("/", DeleteFilmActor)

	return router
}
