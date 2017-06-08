package main

import (
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	placeDetails = "https://maps.googleapis.com/maps/api/place/details/json?&key=%s&placeid=%s"
)

type (
	googleWrap struct {
		Result googlePlace `json:"result"`
		Status string      `json:"status"`
	}

	googlePlace struct {
		PlaceID          string `json:"place_id"`
		Name             string `json:"name"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location coordinates `json:"location"`
		} `json:"geometry"`
		Photos []googlePhoto `json:"photos"`
	}

	googlePhoto struct {
		PhotoReference string `json:"photo_reference"`
	}

	coordinates struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	resultPlace struct {
		PlaceID          string      `json:"place_id"`
		Name             string      `json:"name"`
		FormattedAddress string      `json:"formatted_address"`
		PhotoReference   string      `json:"photo_reference"`
		Coordinates      coordinates `json:"coordinates"`
	}
)

var (
	opts struct {
		Places     []string `short:"p" description:"Google Place Id" env:"CHAPTER1_PLACES" env-delim:"," required:"true"`
		OutputFile string   `short:"o" description:"Output file" env:"CHAPTER1_OUTPUT" required:"true"`
		GoogleKey  string   `short:"k" description:"Google Places API Key" env:"CHAPTER1_GOOGLE_KEY" required:"true"`
	}
)

func main() {
	// read flags, also could be done by `_, err := flags.Parse()`
	if _, err := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash).Parse(); err != nil {
		log.Fatalln(err)
	}

	// get info about the output destination
	fi, err := os.Stat(opts.OutputFile)

	// if the error is "not exists" then we are good
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}

	// unable to write into directory
	if fi.IsDir() {
		log.Fatalln("Output file is a directory")
	}

	var places []resultPlace

	for _, id := range opts.Places {
		log.Printf("Fetching place %q...", id)

		googlePlace, err := fetchPlaceId(id)
		if err != nil {
			log.Println(err)
			continue
		}

		place := resultPlace{
			PlaceID:          googlePlace.PlaceID,
			Name:             googlePlace.Name,
			FormattedAddress: googlePlace.FormattedAddress,
			Coordinates:      googlePlace.Geometry.Location,
		}

		if len(googlePlace.Photos) > 0 {
			place.PhotoReference = googlePlace.Photos[0].PhotoReference
		}

		places = append(places, place)
	}

	// encode slice to json with pretty-print
	data, err := json.MarshalIndent(places, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	// write the whole bunch of data to the file
	if err = ioutil.WriteFile(opts.OutputFile, data, os.ModePerm); err != nil {
		log.Fatalln(err)
	}
}

func fetchPlaceId(id string) (*googlePlace, error) {
	// perform get request
	resp, err := http.Get(fmt.Sprintf(placeDetails, opts.GoogleKey, id))
	if err != nil {
		return nil, err
	}

	// check the status of response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}

	// do not forget to close the body
	defer resp.Body.Close()

	// read all body at once
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// decode to the structure
	var place googleWrap
	if err = json.Unmarshal(data, &place); err != nil {
		return nil, err
	}

	// inside the structure we have status also
	if place.Status != "OK" {
		return nil, fmt.Errorf("Unexpected responce status %q", place.Status)
	}

	// return the only place information
	return &place.Result, nil
}
