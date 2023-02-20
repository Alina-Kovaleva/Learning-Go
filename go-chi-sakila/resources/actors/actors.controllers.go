package actors

import (
	"net/http"

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
	db.DB.First(&actor, id)
	render.Render(w, r, NewActorResponse(&actor))
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	actor := data.Actor

	db.DB.Create(actor)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewActorResponse(actor))
}
