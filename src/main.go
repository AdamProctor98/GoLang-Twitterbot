package main

import (
	"encoding/json"
	"golangtwitterreader/services"
	"log"
	"net/http"
	"os"
)

func searchTweets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response, err := services.PerformSearch(r.URL.Query().Get("query"))

	if err != nil {
		HandleError(w, err)
		return
	}

	json, jsonError := json.Marshal(response)

	if jsonError != nil {
		HandleError(w, jsonError)
		return
	}

	w.Write(json)
	return
}

func HandleError(w http.ResponseWriter, err error) {
	errorJson, jsonError := json.Marshal(struct {
		Error string
	}{err.Error()})

	if jsonError != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", jsonError)
		w.Write([]byte(jsonError.Error()))
	}

	w.Write(errorJson)
}

func main() {

	os.Setenv("BearerToken", "Bearer_Token_Here")

	http.HandleFunc("/tweetbot/", searchTweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
