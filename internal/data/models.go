package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies interface {
		Delete(id int64) error
		Get(id int64) (*Movie, error)
		Insert(movie *Movie) error
		Update(movie *Movie) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}

func NewMockModel() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
