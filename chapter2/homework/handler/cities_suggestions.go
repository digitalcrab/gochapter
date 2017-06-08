package handler

import (
	"gochapter/chapter2/homework/response"
	"googlemaps.github.io/maps"
	"log"
	"net/http"
)

type (
	citiesSuggestionsHandler struct {
		googleMaps *maps.Client
	}

	citiesSuggestionsItem struct {
		PlaceID string `json:"place_id"`
		Name    string `json:"name"`
	}
)

func NewCitiesSuggestions(googleMaps *maps.Client) http.Handler {
	return &citiesSuggestionsHandler{googleMaps}
}

func (h *citiesSuggestionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validate query param
	query := r.URL.Query().Get("q")
	if query == "" {
		response.WriteErrorString(w, http.StatusBadRequest, "Query parameter is missing")
		return
	}

	// build autocomplete request
	req := maps.PlaceAutocompleteRequest{
		Input:    query,
		Types:    maps.AutocompletePlaceTypeCities,
		Language: "en",
	}

	// perform and handle an error
	res, err := h.googleMaps.PlaceAutocomplete(r.Context(), &req)
	if err != nil {
		response.WriteErrorString(w, http.StatusInternalServerError, "Unexpected error on fetching cities")
		log.Printf("CitiesSuggestions: %v\n", err)
		return
	}

	// gather necessary information
	payload := make([]citiesSuggestionsItem, len(res.Predictions))
	for i, p := range res.Predictions {
		payload[i] = citiesSuggestionsItem{
			PlaceID: p.PlaceID,
			Name:    p.Description,
		}
	}

	response.WriteSuccess(w, payload)
}
