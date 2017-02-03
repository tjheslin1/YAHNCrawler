package crawler

import (
	"testing"
	"fmt"
)

func TestRetrievesHackerNewsIdsFromTopStories(t *testing.T) {
	ids := crawlForIntArray("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")

	if(len(ids) == 0) {
		fmt.Println("Expected some ids! Array was empty")
		t.Fail()
	}
}