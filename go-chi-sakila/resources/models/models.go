package models

import "time"

type Actor struct {
	ActorId    int       `gorm:"type:smallint;primaryKey"`
	FirstName  string    `gorm:"type:varchar(45)"`
	LastName   string    `gorm:"type:varchar(45)"`
	LastUpdate time.Time `gorm:"autoCreateTime"`
	Films      []Film    `gorm:"many2many:film_actor;foreignKey:actor_id;joinForeignKey:actor_id;references:film_id;joinReferences:film_id"`
}

type Film struct {
	FilmId      int    `gorm:"column:film_id;primaryKey;autoIncrement"`
	Title       string `gorm:"type:varchar(128);not null"`
	Description string `gorm:"type:text"`
	ReleaseYear int    `gorm:"type:year"`
	LanguageID  int    `gorm:"type:tinyint;not null"`

	RentalDuration  int       `gorm:"type:tinyint;not null;default:3"`
	RentalRate      float64   `gorm:"type:decimal(4,2);not null;default:4.99"`
	Length          int       `gorm:"type:smallint"`
	ReplacementCost float64   `gorm:"type:decimal(5,2);not null;default:19.99"`
	Rating          string    `gorm:"type:enum('G','PG','PG-13','R','NC-17');default:'G'"`
	SpecialFeatures string    `gorm:"type:set('Trailers','Commentaries','Deleted Scenes','Behind the Scenes')"`
	LastUpdate      time.Time `gorm:"autoCreateTime"`
	Actors          []Actor   `gorm:"many2many:film_actor;foreignKey:film_id;joinForeignKey:film_id;references:actor_id;joinReferences:actor_id"`
}

type FilmActor struct {
	ActorID    uint16    `gorm:"primaryKey"`
	FilmID     uint16    `gorm:"primaryKey"`
	LastUpdate time.Time `gorm:"autoCreateTime"`
}

func (Actor) TableName() string {
	return "actor"
}

func (Film) TableName() string {
	return "film"
}

func (FilmActor) TableName() string {
	return "film_actor"
}
