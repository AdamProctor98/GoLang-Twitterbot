package models

type SearchResponse struct {
	Data   []Data `json:"data"`
	Meta   Meta   `json:"meta"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (sr SearchResponse) DataCount() int {
	return len(sr.Data)
}

func (sr SearchResponse) TotalWordCount() int {
	var count int

	for i := 0; i < len(sr.Data); i++ {
		count += sr.Data[i].WordCount()
	}

	return count
}
