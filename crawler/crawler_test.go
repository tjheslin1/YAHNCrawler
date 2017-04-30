package crawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCrawlTopStories(t *testing.T) {
	/*
		serveMux := http.NewServeMux()
		serveMux.Handle("/v0/topstories.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, `[ 14223020 ]`)
		}))

		serveMux.Handle("/v0/item/14223020.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, storyToJSON(exampleStory))
		}))
	*/
	testServer := setUpTestServer([]int{14223020})
	defer testServer.Close()

	logOutput := new(bytes.Buffer)
	logger := log.New(logOutput, "", log.Ldate|log.Ltime)

	topStories := CrawlTopStories(1, testServer.URL, logger)

	if len(topStories) != 1 {
		fmt.Printf("Expected '1' story to be returned but got '%v'\n", len(topStories))
	}

	if !reflect.DeepEqual(*topStories[0], exampleStory) {
		fmt.Printf("Expected \n'%v'\nto equal\n'%v'\n", topStories[0], exampleStory)
		t.Fail()
	}
}

func TestCrawlTopStoriesRetrieveCount(t *testing.T) {

}

func TestQueryIDs(t *testing.T) {
	testIDs := []int{14223020, 14222823, 14219760, 14221229, 14223129, 14221848}
	testServer := setUpTestServer(testIDs)
	defer testServer.Close()

	ids := queryIDs(len(testIDs), testServer.URL+"/v0/topstories.json")

	if len(ids) != len(testIDs) {
		fmt.Printf("Expected '%v' ids! Found '%v' ids\n", len(testIDs), len(ids))
		t.Fail()
	}
}

func TestQueryStory(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, storyToJSON(exampleStory))
	}))
	defer testServer.Close()

	story := queryStory(testServer.URL)

	if !reflect.DeepEqual(*story, exampleStory) {
		fmt.Printf("Expected returned story:\n"+
			"'%v'\n"+
			"to match\n"+
			"example story:\n"+
			"'%v'.", story, exampleStory)
		t.Fail()
	}
}

func setUpTestServer(storyIDs []int) *httptest.Server {
	serveMux := http.NewServeMux()
	serveMux.Handle("/v0/topstories.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, idsToJSON(storyIDs))
	}))

	for storyID := range storyIDs {
		serveMux.Handle("/v0/item/"+string(storyID)+".json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, storyToJSON(exampleStory))
		}))
	}

	return httptest.NewServer(serveMux)
}

func storyToJSON(example Story) string {
	storyJSON, err := json.Marshal(example)
	if err != nil {
		panic(err)
	}

	return string(storyJSON)
}

func idsToJSON(ids []int) string {
	storyJSON, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}

	return string(storyJSON)
}

var exampleStory = Story{
	By:          "prostoalex",
	Descendants: 395,
	ID:          13561388,
	Kids:        []int{13561709, 13561452},
	Score:       506,
	Time:        1486142414,
	Title:       "Amazon soars to more than 341K employees, adding 110K people in a single year",
	Type:        "story",
	URL:         "http://www.geekwire.com/2017/amazon-soars-340k-employees-adding-110k-people-single-year/",
}
