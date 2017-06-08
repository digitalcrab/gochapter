package main

import (
	"github.com/gorilla/mux"
	"gochapter/chapter2/homework"
	"gochapter/chapter2/homework/handler"
	"googlemaps.github.io/maps"
	"log"
	"net/http"
)

func main() {
	// read flags as a config
	config, err := homework.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// initialise google api
	googleMaps, err := maps.NewClient(maps.WithAPIKey(config.GoogleKey))
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.NotFoundHandler = handler.NewNotFound()

	// GET /cities-suggestions
	r.Methods(http.MethodGet).
		Path("/cities-suggestions").
		Handler(handler.NewCitiesSuggestions(googleMaps))

	// POST /cities-info
	r.Methods(http.MethodPost).
		Path("/cities-info").
		Handler(handler.NewCitiesInfo(googleMaps))

	// run HTTP server
	log.Fatalln(http.ListenAndServe(config.HTTPListen, handler.NewRecovery(r)))
}
