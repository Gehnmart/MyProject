package handler

import (
	"MyProject/EasyMessage/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HomeRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user model.User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("user:%+v\n", user)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/home.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}
