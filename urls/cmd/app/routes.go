package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/urls", app.all).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/urls", app.insert).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/urls/{id}", app.get).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/urls", app.deleteAll).Methods(http.MethodDelete)
	r.Use(mux.CORSMethodMiddleware(r))

	return r
}
