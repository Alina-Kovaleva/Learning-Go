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

// CREATE TABLE `film_actor` (
// 	`actor_id` smallint unsigned NOT NULL,
// 	`film_id` smallint unsigned NOT NULL,
// 	`last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`actor_id`,`film_id`),
// 	KEY `idx_fk_film_id` (`film_id`),
// 	CONSTRAINT `fk_film_actor_actor` FOREIGN KEY (`actor_id`) REFERENCES `actor` (`actor_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
// 	CONSTRAINT `fk_film_actor_film` FOREIGN KEY (`film_id`) REFERENCES `film` (`film_id`) ON DELETE RESTRICT ON UPDATE CASCADE
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

//   CREATE TABLE `film` (
// 	`film_id` smallint unsigned NOT NULL AUTO_INCREMENT,
// 	`title` varchar(128) NOT NULL,
// 	`description` text,
// 	`release_year` year DEFAULT NULL,
// 	`language_id` tinyint unsigned NOT NULL,
// 	`original_language_id` tinyint unsigned DEFAULT NULL,
// 	`rental_duration` tinyint unsigned NOT NULL DEFAULT '3',
// 	`rental_rate` decimal(4,2) NOT NULL DEFAULT '4.99',
// 	`length` smallint unsigned DEFAULT NULL,
// 	`replacement_cost` decimal(5,2) NOT NULL DEFAULT '19.99',
// 	`rating` enum('G','PG','PG-13','R','NC-17') DEFAULT 'G',
// 	`special_features` set('Trailers','Commentaries','Deleted Scenes','Behind the Scenes') DEFAULT NULL,
// 	`last_update` datetime(3) DEFAULT NULL,
// 	PRIMARY KEY (`film_id`),
// 	KEY `idx_title` (`title`),
// 	KEY `idx_fk_language_id` (`language_id`),
// 	KEY `idx_fk_original_language_id` (`original_language_id`),
// 	CONSTRAINT `fk_film_language` FOREIGN KEY (`language_id`) REFERENCES `language` (`language_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
// 	CONSTRAINT `fk_film_language_original` FOREIGN KEY (`original_language_id`) REFERENCES `language` (`language_id`) ON DELETE RESTRICT ON UPDATE CASCADE
//   ) ENGINE=InnoDB AUTO_INCREMENT=1018 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

// CREATE TABLE `actor` (
//    `actor_id` smallint unsigned NOT NULL AUTO_INCREMENT,
//    `first_name` varchar(45) NOT NULL,
//    `last_name` varchar(45) NOT NULL,
//    `last_update` datetime(3) DEFAULT NULL,
//    PRIMARY KEY (`actor_id`),
//    KEY `idx_actor_last_name` (`last_name`)
//  ) ENGINE=InnoDB AUTO_INCREMENT=211 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
