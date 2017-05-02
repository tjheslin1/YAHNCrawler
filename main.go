package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tjheslin1/YAHNCrawler/crawler"
	"github.com/tjheslin1/YAHNCrawler/html"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	fmt.Println("Yet Another Hacker News Crawler!")
	stories := crawler.CrawlTopStories(10, "https://hacker-news.firebaseio.com", logger)

	storyContext := html.StoriesContext{
		Title:   "Hacker News crawler",
		Stories: stories,
	}

	html.GenerateHTML(storyContext, os.Stdout)

	var response int
	fmt.Scanf("%c", &response)
}
