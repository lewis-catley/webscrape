package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
	"github.com/lewis-catley/webscrape/urls/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	// Get the urls
	urls, err := app.urls.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Convert bson to json
	b, err := json.Marshal(urls)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}
	params := mux.Vars(r)
	u, err := app.urls.GetByID(params["id"])
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Convert bson to json
	b, err := json.Marshal(u)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	url := &models.URLPost{}
	err := json.NewDecoder(r.Body).Decode(url)
	if err != nil {
		app.serverError(w, err) // TODO: This is probably not a 500, but a problem with the request
		return
	}

	ir, err := app.urls.Insert(*url)
	if err != nil {
		app.serverError(w, err)
		return
	}

	b, err := json.Marshal(models.URLMessage{
		ID:  ir.InsertedID.(primitive.ObjectID).Hex(),
		URL: url.URL,
	})
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Fire the request to the web scraper application
	_, err = http.Post("http://scraper:4001/urls", "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Error sending request to web scraper", err)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ir.InsertedID.(primitive.ObjectID).Hex()))
}

func (app *application) deleteAll(w http.ResponseWriter, r *http.Request) {
	err := app.urls.DeleteAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNoContent)
}
