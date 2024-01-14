package main

import (
	"MyProject/EasyMessage/handler"
	"net/http"
)

func handlefunc() {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup/", handler.SignupHandler)
	mux.HandleFunc("/register/", handler.RegisterRequest)

	mux.HandleFunc("/signin/", handler.SigninHandler)
	mux.HandleFunc("/login/", handler.LoginRequest)
	mux.HandleFunc("/home/", handler.HomeHandler)

	http.ListenAndServe(":8080", mux)
}

func main() {
	handlefunc()
}
