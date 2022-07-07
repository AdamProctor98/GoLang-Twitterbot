package models

import "strings"

type Data struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func (d Data) WordCount() int {
	sanitizedTweet := strings.Replace(d.Text, "\n", " ", -1)
	return len(strings.Split(sanitizedTweet, " "))
}
