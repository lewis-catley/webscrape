package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/lewis-catley/webscrape/scraper/pkg/models"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) trigger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	up := &models.URLPost{}
	err := json.NewDecoder(r.Body).Decode(up)
	if err != nil {
		app.serverError(w, err) // TODO: This is probably not a 500, but a problem with the request
		return
	}

	app.triggers <- up
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNoContent)
	return
}
