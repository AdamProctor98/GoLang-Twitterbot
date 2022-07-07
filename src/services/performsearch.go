package services

import (
	"errors"
	"golangtwitterreader/models"
)

func PerformSearch(query string) (*models.SearchSummary, error) {

	if query == "" {
		return nil, errors.New("Invalid Request: you must provide a valid query")
	}

	responses, err := SearchTweets(query, "")

	if err != nil {
		return nil, err
	}

	var totalResultCount int
	var totalWordCount int

	for i := 0; i < len(responses); i++ {
		response := responses[i]

		totalResultCount += response.DataCount()
		totalWordCount += response.TotalWordCount()
	}

	avgWordCount := totalWordCount / totalResultCount

	summary := &models.SearchSummary{SearchTerm: query, ResultCount: totalResultCount, AvgWordCount: avgWordCount}

	return summary, nil
}
