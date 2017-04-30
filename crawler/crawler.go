package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

const storyCountLimit = 500

// CrawlTopStories retrieves the top 500 Hacker News articles and retrieves
// their Story information. This is presented in HTML.
func CrawlTopStories(maxStoryCount int, hostname string, logger *log.Logger) []*Story {
	start := time.Now()
	topStoryIDs := queryIDs(maxStoryCount, hostname+"/v0/topstories.json")

	storyChan := make(chan *Story, len(topStoryIDs))
	go func(ids []int, storyCh chan<- *Story) {
		for _, topStoryID := range ids {
			storyCh <- queryStory(hostname + "/v0/item/" + strconv.Itoa(topStoryID) + ".json")
			time.Sleep(50 * time.Millisecond)
		}
		close(storyChan)
	}(topStoryIDs, storyChan)

	var topStories []*Story
	for story := range storyChan {
		topStories = append(topStories, story)
	}

	logger.Printf("Querying stories took '%v'", time.Since(start))

	return topStories
}

func queryIDs(maxStoryCount int, url string) []int {
	var count int
	if maxStoryCount < storyCountLimit {
		count = maxStoryCount
	} else {
		count = storyCountLimit
	}

	ids := make([]int, count)
	query(url, &ids)
	return ids
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
	}
}
