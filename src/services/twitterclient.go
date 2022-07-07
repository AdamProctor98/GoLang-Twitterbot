package services

import (
	"encoding/json"
	"fmt"
	"golangtwitterreader/models"
	"io/ioutil"
	"net/http"
)

func SearchRecentTweets(url string, bearerToken string) (*models.SearchResponse, error) {
	req, reqErr := http.NewRequest("GET", url, nil)

	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Add("Authorization", "Bearer "+bearerToken)
	client := &http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to search recent tweets. Status: %s", resp.Status)
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return nil, readErr
	}

	var response models.SearchResponse

	jsonerr := json.Unmarshal(body, &response)

	if jsonerr != nil {
		return nil, jsonerr
	}

	if response.Status != 0 {

		return nil, fmt.Errorf("Status Code: %d. Details: %s", response.Status, response.Detail)
	}

	return &response, nil
}
