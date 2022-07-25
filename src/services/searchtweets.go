package services

import (
	"errors"
	"golangtwitterreader/models"
	"os"
)

func SearchTweets(query string, nextToken string) ([]*models.SearchResponse, error) {
	bearerToken, ok := os.LookupEnv("BearerToken")

	if ok == false {
		return nil, errors.New("Invalid Request. Bearer Token is not set")
	}

	url := "https://api.twitter.com/2/tweets/search/recent?query=" + query + "&max_results=100"

	if nextToken != "" {
		url = url + "&next_token=" + nextToken
	}

	var responses []*models.SearchResponse

	response, resErr := SearchRecentTweets(url, bearerToken)

	if resErr != nil {
		return nil, resErr
	}

	responses = append(responses, response)

	if response.Meta.NextToken != "" {
		result, err := SearchTweets(query, response.Meta.NextToken)

		if err != nil {
			return nil, err
		}

		responses = append([]*models.SearchResponse(responses), []*models.SearchResponse(result)...)
	}

	return responses, nil
}
