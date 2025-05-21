package main

import (
	"fmt"
	"net/http"
	"time"

	"bookliba.askvart.ru/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create new film")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.ReadIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	movie := data.Movie{
		ID: id,
		CreatedAt: time.Now(),
		Title: "The Matrix",
		Year: 1999,
		Runtime: 121,
		Genres: []string {"fantastic", "war", "triller"},
		Version: 1,
	}
	err = app.WriteJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
