package zipit

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// ZippoPlace describe the place in the zippopatomus response
type ZippoPlace struct {
	City      string `json:"place name"`
	Longitude string `json:"longitude"`
	State     string `json:"state"`
	StateCode string `json:"state abbreviation"`
	Latitude  string `json:"latitude"`
}

// ZippoResponse describes the response from zippopatomus
type ZippoResponse struct {
	PostCode            string       `json:"post code"`
	Country             string       `json:"country"`
	CountryAbbreviation string       `json:"country abbreviation"`
	Places              []ZippoPlace `json:"places"`
}

//NewZippoResponse takes the content of the response and returns a parsed ZippoResponse object
func NewZippoResponse(content io.Reader) (ZippoResponse, error) {
	zippoResponse := ZippoResponse{}
	contentBytes, err := ioutil.ReadAll(content)
	if err != nil {
		return zippoResponse, fmt.Errorf("error reading content: %s", err.Error())
	}
	err = json.Unmarshal(contentBytes, &zippoResponse)
	if err != nil {
		return zippoResponse, fmt.Errorf("error unmarshalling content: %s", err.Error())
	}
	return zippoResponse, nil
}

// GetDetailsFor gets the details for the zipcode from zippopotam.us and returns a ZippoResponse
func GetDetailsFor(zipcode string, client *http.Client) (ZippoResponse, error) {
	url := fmt.Sprintf("http://api.zippopotam.us/us/%s", zipcode)
	resp, err := client.Get(url)
	if err != nil {
		return ZippoResponse{}, fmt.Errorf("error requesting from zippo: %s", err)
	}
	if resp.StatusCode != 200 {
		return ZippoResponse{}, fmt.Errorf("zippopotam.us returned %d status code for %s", resp.StatusCode, url)
	}
	return NewZippoResponse(resp.Body)
}

//GetPlaceFor returns the place given by a zipcode
func GetPlaceFor(zipcode string, client *http.Client) (ZippoPlace, error) {
	zippo, err := GetDetailsFor(zipcode, client)
	if err != nil {
		return ZippoPlace{}, err
	}
	if len(zippo.Places) != 1 {
		return ZippoPlace{}, fmt.Errorf("error in zippo response: zipcode returned more than one city for zipcode: %s", zipcode)
	}
	return zippo.Places[0], nil
}
