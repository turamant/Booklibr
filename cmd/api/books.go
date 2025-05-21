package main

import (
	"fmt"
	"net/http"
	"time"

	"booklibr.askvart.ru/internal/data"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create new book")
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.ReadIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	book := data.Book{
		ID: id,
		CreatedAt: time.Now(),
		Title: "The war and piece",
		Year: 1999,
		Pages: 121,
		Genres: []string {"fantastic", "war", "triller"},
		Author: "Tolstoy L.N.",
	}
	err = app.WriteJson(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
