package html

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/tjheslin1/YAHNCrawler/crawler"
)

func TestGenerateHTML(t *testing.T) {
	storiesContext := StoriesContext{
		Title:   "TestTitle",
		Stories: []*crawler.Story{&exampleStory},
	}

	var expectedOuput = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>TestTitle</title>
</head>
<body>
    <div>prostoalex</div>
</body>
</html>`

	output := new(bytes.Buffer)
	GenerateHTML(storiesContext, output)

	actual := string(output.Bytes())
	if actual != expectedOuput {
		fmt.Printf("Expected '%v'\nto equal\n'%v'\n", actual, expectedOuput)
		t.Fail()
	}
}

var exampleStory = crawler.Story{
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
