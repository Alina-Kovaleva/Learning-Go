package films

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"tsi.co/go-chi-sakila/resources/models"
)

type FilmRequest struct {
	*models.Film
}

func (f *FilmRequest) Bind(r *http.Request) error {
	if f.Film == nil {
		return errors.New("missing required Film fields")
	}

	return nil
}

type FilmResponse struct {
	*models.Film
}

func NewFilmResponse(film *models.Film) *FilmResponse {
	return &FilmResponse{film}
}

func NewFilmListResponse(films []*models.Film) []render.Renderer {
	list := []render.Renderer{}
	for _, film := range films {
		list = append(list, NewFilmResponse(film))
	}
	return list
}

func (f *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
