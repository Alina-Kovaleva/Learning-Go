package actors

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"tsi.co/go-chi-sakila/resources/models"
)

type ActorUpdate struct {
	FirstName *string
	LastName  *string
}

type ActorRequest struct {
	*models.Actor
}

type ActorUpdateRequest struct {
	*ActorUpdate
}

func (a *ActorRequest) Bind(r *http.Request) error {
	if a.Actor == nil {
		return errors.New("missing required Actor fields")
	}

	a.Actor.FirstName = strings.ToUpper(a.Actor.FirstName)
	a.Actor.LastName = strings.ToUpper(a.Actor.LastName)

	return nil
}

func (a *ActorUpdateRequest) Bind(r *http.Request) error {
	if a.ActorUpdate == nil {
		return errors.New("missing required Actor fields")
	}
	return nil
}

type ActorResponse struct {
	*models.Actor
}

func NewActorResponse(actor *models.Actor) *ActorResponse {
	return &ActorResponse{actor}
}

func NewActorListResponse(actors []*models.Actor) []render.Renderer {
	list := []render.Renderer{}
	for _, actor := range actors {
		list = append(list, NewActorResponse(actor))
	}
	return list
}

func (a *ActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
