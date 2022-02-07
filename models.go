package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Movie struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"Description"`
	Year        int            `json:"Year"`
	ReleaseDate time.Time      `json:"release_date"`
	Runtime     int            `json:"runtime"`
	Rating      int            `json:"rating"`
	CreatedAt   time.Time      `json:"_"`
	UpdatedAt   time.Time      `json:"_"`
	MovieGenre  map[int]string `json:"genres"`
}

type Genre struct {
	ID        int       `json:"_"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}

type MovieGenre struct {
	ID        int       `json:"_"`
	MovieID   int       `json:"_"`
	GenerID   int       `json:"_"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}
