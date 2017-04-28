package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Story represents a story on Hacker News.
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

func CrawlTopStories() {
	topStoryIDs := queryIDs("https://hacker-news.firebaseio.com/v0/topstories.json")

	for _, topStoryID := range *topStoryIDs {
		fmt.Println(queryStory("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(topStoryID) + ".json?print=pretty"))
	}
}

// ("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
// https://hacker-news.firebaseio.com/v0/item/8863.json?print=pretty
func queryIDs(url string) *[]int {
	ids := make([]int, 500)
	query(url, &ids)
	return &ids
}

func queryStory(url string) *Story {
	story := new(Story)
	query(url, story)
	return story
}

func query(url string, result interface{}) *interface{} {
	resp, err := http.Get(url)
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	unmarshalErr := json.Unmarshal(body, result)
	check(unmarshalErr)

	return &result
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
