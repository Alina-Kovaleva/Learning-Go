package filmactor

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"tsi.co/go-chi-sakila/resources/models"
)

type FilmActorRequest struct {
	*models.FilmActor
}

func (f *FilmActorRequest) Bind(r *http.Request) error {
	if f.FilmActor == nil {
		return errors.New("missing required FilmActor fields")
	}

	return nil
}

type FilmActorResponse struct {
	*models.FilmActor
}

func NewFilmActorResponse(fa *models.FilmActor) *FilmActorResponse {
	return &FilmActorResponse{fa}
}

func NewFilmActorListResponse(fas []*models.FilmActor) []render.Renderer {
	list := []render.Renderer{}
	for _, fa := range fas {
		list = append(list, NewFilmActorResponse(fa))
	}
	return list
}

func (f *FilmActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
