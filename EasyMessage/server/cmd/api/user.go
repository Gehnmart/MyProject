package main

import (
	data "MyProject/EasyMessage/server/internal"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelop{
		"status": "available",
		"system_info": map[string]string{
			"enviroment": app.config.env,
			"version":    verson,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponce(w, r, err)
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CreateNewMovieHandler")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil || id < 1 {
		app.notFoundResponce(w, r)
		return
	}

	user := data.User{
		Id:        id,
		Name:      "Ivan",
		ShortName: "ivan",
		Email:     "goi234@gmail.com",
		Password:  "jaskldfjalksjdfl",
		Rooms:     nil,
	}
	err = app.writeJSON(w, http.StatusOK, envelop{"user": user}, nil)
	if err != nil {
		app.serverErrorResponce(w, r, err)
	}
}

func (app *application) createNewUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	fmt.Fprintf(w, "%+v", input)
}

func (app *application) readIdParam(r *http.Request) (uint64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, err
	}

	return id, err
}
