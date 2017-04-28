package crawler

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

// ("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
func crawlForIntArray(url string) []int {
	resp, err := http.Get(url)
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	newsIds := make([]int, 500)
	unmarshalErr := json.Unmarshal(body, &newsIds)
	check(unmarshalErr)

	return newsIds
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}