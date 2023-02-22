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
	"tsi.co/go-chi-sakila/resources/models"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	var films []*models.Film
	db.DB.Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var film models.Film
	result := db.DB.Model(&models.Film{}).Preload("Actors").First(&film, id)
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
		return
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

	var film models.Film
	result := db.DB.First(&film, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	var filmActors []models.FilmActor
	result = db.DB.Where("film_id = ?", id).Find(&filmActors)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	result = db.DB.Delete(&filmActors)
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
		return
	}

	film := data.Film
	id := chi.URLParam(r, "id")

	var updatedFilm models.Film
	result := db.DB.First(&updatedFilm, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	if film.Title != "" {
		updatedFilm.Title = film.Title
		updatedFilm.Description = film.Description
		updatedFilm.ReleaseYear = film.ReleaseYear
		updatedFilm.LanguageID = film.LanguageID
		updatedFilm.RentalDuration = film.RentalDuration
		updatedFilm.RentalRate = film.RentalRate
		updatedFilm.Length = film.Length
		updatedFilm.ReplacementCost = film.ReplacementCost
		updatedFilm.Rating = film.Rating
		updatedFilm.SpecialFeatures = film.SpecialFeatures
	}

	updatedFilm.LastUpdate = time.Now()
	db.DB.Save(&updatedFilm)

	log.Println("Actor's info was updated")
	render.Render(w, r, NewFilmResponse(&updatedFilm))
	render.Status(r, http.StatusNoContent)
}

func GetFilmsByActorId(w http.ResponseWriter, r *http.Request) {
	actorID := chi.URLParam(r, "id")

	var filmActors []models.FilmActor
	result := db.DB.Where("actor_id = ?", actorID).Find(&filmActors)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	var films []*models.Film
	for _, fa := range filmActors {
		var film models.Film
		result := db.DB.First(&film, fa.FilmID)
		if result.Error != nil {
			render.Render(w, r, e.ErrInvalidRequest(result.Error))
			return
		}

		films = append(films, &film)
	}

	render.RenderList(w, r, NewFilmListResponse(films))
}
