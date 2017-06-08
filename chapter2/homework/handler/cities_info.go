package handler

import (
	"gochapter/chapter2/homework/request"
	"gochapter/chapter2/homework/response"
	"googlemaps.github.io/maps"
	"log"
	"net/http"
)

type (
	citiesInfoHandler struct {
		googleMaps *maps.Client
	}

	citiesInfoItem struct {
		PlaceID          string `json:"place_id"`
		Name             string `json:"name"`
		FormattedAddress string `json:"formatted_address"`
		PhotoReference   string `json:"photo_reference,omitempty"`
		Coordinates      struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"coordinates"`
	}
)

func NewCitiesInfo(googleMaps *maps.Client) http.Handler {
	return &citiesInfoHandler{googleMaps}
}

func (h *citiesInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read POST request and handle an error
	var in []string
	if err := request.ParseBody(r.Body, &in); err != nil {
		response.WriteErrorString(w, http.StatusBadRequest, err.Error())
		return
	}

	// nothing to do
	if len(in) == 0 {
		response.WriteErrorString(w, http.StatusBadRequest, "You must specify at least one place id")
		return
	}

	var payload []citiesInfoItem

	for _, id := range in {
		// build request
		req := maps.PlaceDetailsRequest{
			PlaceID:  id,
			Language: "en",
		}

		// perform and handle an error
		res, err := h.googleMaps.PlaceDetails(r.Context(), &req)
		if err != nil {
			log.Printf("Error on get place %q details: %v\n", id, err)
			continue
		}

		item := citiesInfoItem{
			PlaceID:          res.PlaceID,
			Name:             res.Name,
			FormattedAddress: res.FormattedAddress,
			Coordinates:      res.Geometry.Location,
		}

		if len(res.Photos) > 0 {
			item.PhotoReference = res.Photos[0].PhotoReference
		}

		payload = append(payload, item)
	}

	response.WriteSuccess(w, payload)
}
