package main

import (
	"fmt"

	"github.com/tjheslin1/YAHNCrawler/crawler"
)

func main() {
	fmt.Println("Yet Another Hacker News Crawler!")
	crawler.CrawlTopStories("https://hacker-news.firebaseio.com")

	var response int
	fmt.Scanf("%c", &response)
}
