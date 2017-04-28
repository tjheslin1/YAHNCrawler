package crawler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestQueryIDs(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[ 14223020, 14222823, 14219760, 14221229, 14223129, 14221848 ]`)
	}))
	defer testServer.Close()

	ids := queryIDs(testServer.URL)

	if len(*ids) != 6 {
		fmt.Printf("Expected 6 ids! Found '%v' ids\n", len(*ids))
		t.Fail()
	}
}

func TestQueryStory(t *testing.T) {
	storyJSON, err := json.Marshal(exampleStory)
	if err != nil {
		fmt.Printf("Error marshalling expected json. '%s'.\n", err)
		t.Fail()
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(storyJSON))
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
