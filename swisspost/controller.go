package swisspost

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// getURL build the url
func (s Swisspost) getURL(bfsNr int) (url string) {
	url = apiURL + dataset +
		"?" + dSelect +
		"&" + dWhere + fmt.Sprint(bfsNr) +
		"&" + dLimit +
		"&" + dOffset +
		"&" + dTZ +
		"&" + dKey + s.apiKey

	return
}

// GetTown get the city & canton rom the swisspost dataset according to the given bfsNr
func (s Swisspost) GetTown(bfsNr int) (string, string, error) {
	url := s.getURL(bfsNr)

	// Send the request to swisspost with 3 retry in error case
	resp, err := resty.New().SetRetryCount(3).R().Get(url)
	if err != nil {
		return "", "", err
	}

	// Read the response
	var response Response

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return "", "", errors.New("could not unmarshal the JSON structure")
	}

	// Check the response
	if len(response.Records) == 0 {
		return "", "", errors.New("no data found")
	}
	f := response.Records[0].Record.Fields

	return f.Kanton, f.GemeindeName, nil
}
