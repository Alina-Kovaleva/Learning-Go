package actors

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

func ListActors(w http.ResponseWriter, r *http.Request) {
	var actors []*Actor
	db.DB.Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func GetActorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var actor Actor

	result := db.DB.First(&actor, id)

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

	var actor Actor
	result := db.DB.First(&actor, id)

	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	db.DB.Delete(&actor)
	log.Println("Actor was delited by id: ", id)
	render.Status(r, http.StatusNoContent)
}

func UpdateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}
	actor := data.Actor

	id := chi.URLParam(r, "id")

	var updatedActor Actor

	result := db.DB.First(&updatedActor, id)
	if result.Error != nil {
		render.Render(w, r, e.ErrInvalidRequest(result.Error))
		return
	}

	if actor.FirstName == "" && actor.LastName == "" {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("either First Name or Last Name is required")))
		return
	}

	updatedActor.FirstName = actor.FirstName
	updatedActor.LastName = actor.LastName
	updatedActor.LastUpdate = time.Now()

	db.DB.Save(&updatedActor)

	log.Println("Actor's info was updated")
	render.Render(w, r, NewActorResponse(&updatedActor))
	render.Status(r, http.StatusNoContent)
}
