package zipit

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestNewZippopatomus(t *testing.T) {
	sampleResponse := `{
		"post code": "90210",
		"country": "United States",
		"country abbreviation": "US",
		"places": [
			{
				"place name": "Beverly Hills",
				"longitude": "-118.4065",
				"state": "California",
				"state abbreviation": "CA",
				"latitude": "34.0901"
			}
		]
	}`
	zippo, err := NewZippoResponse(strings.NewReader(sampleResponse))
	if err != nil {
		t.Errorf("New Zippo unmarshalling failed becasue of error: %s", err)
	}
	if zippo.Country != "United States" {
		t.Errorf("zippo.Country != 'United States', got \"%s\" instead", zippo.Country)
	}

	if zippo.CountryAbbreviation != "US" {
		t.Errorf("zippo.CountryAbbreviation != \"US\", got \"%s\" instead", zippo.CountryAbbreviation)
	}

	if len(zippo.Places) != 1 {
		t.Errorf("zippo.Places is not equal to 1. Got %d instead", len(zippo.Places))
	}
}

func TestGetDetails(t *testing.T) {
	sampleResponse := `{
		"post code": "90210",
		"country": "United States",
		"country abbreviation": "US",
		"places": [
			{
				"place name": "Beverly Hills",
				"longitude": "-118.4065",
				"state": "California",
				"state abbreviation": "CA",
				"latitude": "34.0901"
			}
		]
	}`
	realZippo, err := NewZippoResponse(strings.NewReader(sampleResponse))
	if err != nil {
		t.Fatal(err)
	}
	zippo, err := GetDetailsFor("90210", &http.Client{})
	if err != nil {
		t.Errorf("Error is not nil: %s", err)
	}

	if !reflect.DeepEqual(realZippo, zippo) {
		t.Errorf("zippo != realZippo: %v", zippo)
	}

	zippo, err = GetDetailsFor("90210hh", &http.Client{})
	if err == nil {
		t.Error("Error is nil")
	}
}

func TestGetCityStateFromZippo(t *testing.T) {
	place, err := GetPlaceFor("90210", &http.Client{})
	if err != nil {
		t.Error(err)
	}

	if place.City != "Beverly Hills" {
		t.Errorf("Expected \"Beverly Hills\", got %v instead", place.City)
	}

	if place.StateCode != "CA" {
		t.Errorf("Expected \"CA\", got %v instead", place.StateCode)
	}
}
