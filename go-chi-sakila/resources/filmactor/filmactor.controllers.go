package filmactor

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	db "tsi.co/go-chi-sakila/database"
	e "tsi.co/go-chi-sakila/error"
	"tsi.co/go-chi-sakila/resources/models"
)

func CreateFilmActor(w http.ResponseWriter, r *http.Request) {
	var data FilmActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}

	filmActor := data.FilmActor

	var actor models.Actor
	if result := db.DB.First(&actor, filmActor.ActorID); result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("actor not found")))
		return
	}

	var film models.Film
	if result := db.DB.First(&film, filmActor.FilmID); result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("film not found")))
		return
	}

	db.DB.Create(filmActor)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFilmActorResponse(filmActor))
}

func DeleteFilmActor(w http.ResponseWriter, r *http.Request) {
	var data FilmActorRequest
	if err := render.Bind(r, &data); err != nil {
		fmt.Println(err)
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}
	filmActor := data.FilmActor

	if filmActor.FilmID == 0 || filmActor.ActorID == 0 {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("both FilmID and ActorID are required")))
		return
	}

	var existingFilmActor models.FilmActor
	if result := db.DB.Where("film_id = ? AND actor_id = ?", filmActor.FilmID, filmActor.ActorID).First(&existingFilmActor); result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("filmActor not found")))
		return
	}

	db.DB.Delete(&existingFilmActor)
	render.Status(r, http.StatusNoContent)
}
