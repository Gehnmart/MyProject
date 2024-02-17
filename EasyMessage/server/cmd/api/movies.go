package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status available\n")
	fmt.Fprintf(w, "enviroment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", verson)
}

func (app *application) createMovieHandler(w http.ResponseWriter, r http.Request) {
	fmt.Fprintln(w, "CreateNewMovieHandler")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil {
		fmt.Fprintln(w, "id movie = %d", id)
	} else {
		fmt.Fprintln(w, "id not found")
	}
}
