package films

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListFilms)
	router.Post("/", CreateFilm)

	router.Get("/{id}", GetFilmById)
	router.Delete("/{id}", DeleteFilm)
	router.Patch("/{id}", UpdateFilm)

	router.Get("/actors/{id}", GetFilmsByActorId)

	return router
}
