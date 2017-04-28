package main

import (
	"fmt"

	"github.com/tjheslin1/YAHNCrawler/crawler"
)

func main() {
	fmt.Println("Yet Another Hacker News Crawler!")
	crawler.CrawlTopStories()
}
