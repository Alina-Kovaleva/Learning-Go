package actors

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	db "tsi.co/go-chi-sakila/database"
	e "tsi.co/go-chi-sakila/error"
	"tsi.co/go-chi-sakila/resources/models"
)

func ListActors(w http.ResponseWriter, r *http.Request) {
	var actors []*models.Actor
	db.DB.Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func GetActorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var actor models.Actor

	result := db.DB.Model(&models.Actor{}).Preload("Films").First(&actor, id)

	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	render.Render(w, r, NewActorResponse(&actor))
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}

	actor := data.Actor

	if actor.FirstName == "" && actor.LastName == "" {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("either First Name or Last Name is required")))
		return
	}

	db.DB.Create(actor)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewActorResponse(actor))
}

func DeleteActor(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var actor models.Actor
	result := db.DB.First(&actor, id)

	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	var filmActors []models.FilmActor
	result = db.DB.Where("actor_id = ?", id).Find(&filmActors)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	if len(filmActors) != 0 {
		result = db.DB.Delete(&filmActors)
		if result.Error != nil {
			render.Render(w, r, e.ErrInvalidRequest(result.Error))
			return
		}
	}

	db.DB.Delete(&actor)
	log.Println("Actor was delited by id: ", id)
	render.Status(r, http.StatusNoContent)
}

func UpdateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorUpdateRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}

	actorUpdate := data.ActorUpdate

	log.Println(actorUpdate)

	if actorUpdate.FirstName == nil && actorUpdate.LastName == nil {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("either First Name or Last Name is required")))
		return
	}

	if actorUpdate.FirstName != nil && actorUpdate.LastName != nil {
		if *actorUpdate.FirstName == "" && *actorUpdate.LastName == "" {
			render.Render(w, r, e.ErrInvalidRequest(errors.New("either First Name or Last Name is required")))
			return
		}

	}

	id := chi.URLParam(r, "id")

	var originalActor models.Actor

	result := db.DB.First(&originalActor, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	if actorUpdate.FirstName != nil {
		originalActor.FirstName = strings.ToUpper(*actorUpdate.FirstName)
	}

	if actorUpdate.LastName != nil {
		originalActor.LastName = strings.ToUpper(*actorUpdate.LastName)
	}

	if originalActor.FirstName == "" && originalActor.LastName == "" {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("either First Name or Last Name is required")))
		return
	}

	originalActor.LastUpdate = time.Now()

	db.DB.Save(&originalActor)

	log.Println("Actor's info was updated")
	render.Render(w, r, NewActorResponse(&originalActor))
	render.Status(r, http.StatusNoContent)
}
