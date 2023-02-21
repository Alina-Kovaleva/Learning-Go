package films

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	db "tsi.co/go-chi-sakila/database"
	e "tsi.co/go-chi-sakila/error"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	var films []*Film
	db.DB.Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var film Film
	result := db.DB.First(&film, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	render.Render(w, r, NewFilmResponse(&film))
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	film := data.Film

	if film.Title == "" && (film.LanguageID < 1 && film.LanguageID > 6) {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("title and language id are required")))
		return
	}

	db.DB.Create(film)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFilmResponse(film))
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var film Film
	result := db.DB.First(&film, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	db.DB.Delete(&film)
	log.Println("Film was delited by id: ", id)
	render.Status(r, http.StatusNoContent)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}
	film := data.Film

	id := chi.URLParam(r, "id")

	var updatedFilm Film

	result := db.DB.First(&updatedFilm, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	if film.Title != "" {
		updatedFilm.Title = film.Title
	}

	updatedFilm.LastUpdate = time.Now()
	db.DB.Save(&updatedFilm)

	log.Println("Actor's info was updated")
	render.Render(w, r, NewFilmResponse(&updatedFilm))
	render.Status(r, http.StatusNoContent)
}
