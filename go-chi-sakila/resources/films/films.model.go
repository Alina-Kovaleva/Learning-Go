package films

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Film struct {
	FilmId             int       `gorm:"column:film_id;primaryKey;autoIncrement"`
	Title              string    `gorm:"type:varchar(128);not null"`
	Description        string    `gorm:"type:text"`
	ReleaseYear        int       `gorm:"type:year"`
	LanguageID         int       `gorm:"type:tinyint;not null"`
	OriginalLanguageId int       `gorm:"type:tinyint;default:null"`
	RentalDuration     int       `gorm:"type:tinyint;not null;default:3"`
	RentalRate         float64   `gorm:"type:decimal(4,2);not null;default:4.99"`
	Length             int       `gorm:"type:smallint"`
	ReplacementCost    float64   `gorm:"type:decimal(5,2);not null;default:19.99"`
	Rating             string    `gorm:"type:enum('G','PG','PG-13','R','NC-17');default:'G'"`
	SpecialFeatures    string    `gorm:"type:set('Trailers','Commentaries','Deleted Scenes','Behind the Scenes')"`
	LastUpdate         time.Time `gorm:"autoCreateTime"`
}

func (Film) TableName() string {
	return "film"
}

type FilmRequest struct {
	*Film
}

func (f *FilmRequest) Bind(r *http.Request) error {
	if f.Film == nil {
		return errors.New("missing required Film fields")
	}

	return nil
}

type FilmResponse struct {
	*Film
}

func NewFilmResponse(film *Film) *FilmResponse {
	return &FilmResponse{film}
}

func NewFilmListResponse(films []*Film) []render.Renderer {
	list := []render.Renderer{}
	for _, film := range films {
		list = append(list, NewFilmResponse(film))
	}
	return list
}

func (f *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
