package main

import (
	"fmt"
	"log"
	"os"
	"path"

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

	wd, _ := os.Getwd()
	file := templateFile()
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	logger.Printf("Template file written to: '%v'\n", path.Join(wd, file.Name()))

	html.GenerateHTML(storyContext, file)
}

func templateFile() *os.File {
	os.Mkdir("out", 0777)
	file, err := os.OpenFile("out/template.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return file
}
