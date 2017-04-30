package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
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

const topStoryCount = 500

// CrawlTopStories TODO
func CrawlTopStories(hostname string) {
	topStoryIDs := queryIDs(hostname + "/v0/topstories.json")

	storyChan := make(chan *Story, topStoryCount)
	go func(storyCh chan<- *Story) {
		for _, topStoryID := range *topStoryIDs {
			// fmt.Println("query prepared")
			storyCh <- queryStory("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(topStoryID) + ".json?print=pretty")
			// fmt.Println("que ry sent")
			time.Sleep(500 * time.Millisecond)
		}
		close(storyChan)
	}(storyChan)

	for story := range storyChan {
		fmt.Println(story)
	}
}

func queryIDs(url string) *[]int {
	ids := make([]int, topStoryCount)
	query(url, &ids)
	return &ids
}

func queryStory(url string) *Story {
	story := new(Story)
	query(url, story)
	return story
}

// query performs a HTTP: GET request to the specified url.
// Storing the response's body in the `result`.
// query expects the provided `result` to conform to the query's repsonse body.
func query(url string, result interface{}) *interface{} {
	resp, err := http.Get(url)
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
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
