package crawler

import "encoding/json"

type Story struct {
	By string
	Descendants int
	Id int
	Kids []int
	Score int
	Time int
	Title string
	Type string
	Url string
}

func UnmarshalJSON(buf []byte) (*Story, error) {
	story := Story{}
	if err := json.Unmarshal(buf, &story); err != nil {
		return nil, err
	}

	return &story, nil
}