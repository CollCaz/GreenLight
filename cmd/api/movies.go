package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CollCaz/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creat new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		if app.config.env == "development" {
			app.logger.Warn(err)
		}
		http.NotFound(w, r)
		return
	}

	data := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "one flew over the cuckoos nest",
		Runtime:   105,
		Genres:    []string{"drama", "psychological"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": data}, nil)
	if err != nil {
		app.logger.Error(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
