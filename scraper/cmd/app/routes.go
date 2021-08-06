package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/urls", app.trigger).Methods(http.MethodPost)
	r.PathPrefix("/debug/").Handler(http.DefaultServeMux)
	r.Use(mux.CORSMethodMiddleware(r))

	return r
}
