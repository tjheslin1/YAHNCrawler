package crawler

import (
	"testing"
	"fmt"
)

func TestUnmarshallsStoryJson(t *testing.T) {
	testJson := "{" +
		"\"by\": \"prostoalex\"," +
		"\"descendants\": 272," +
		"\"id\": 13561388," +
		"\"kids\": [" +
			"13561709," +
			"13561527" +
		"]," +
		"\"score\": 324," +
		"\"time\": 1486142414," +
		"\"title\": \"Amazon soars to more than 341K employees, adding 110K people in a single year\"," +
		"\"type\": \"story\"," +
		"\"url\": \"http://www.geekwire.com/2017/amazon-soars-340k-employees-adding-110k-people-single-year/\"" +
		"}"
	jsonBody := []byte(testJson)

	story, err := UnmarshalJSON(jsonBody)
	if err != nil {
		fmt.Print(err)
		t.Fail()
	}

	if story.By != "prostoalex" {
		fmt.Printf("Expected story.By to equal '%s' but was '%s'\n", "prostoalex", story.By)
		t.Fail()
	}
	if story.Descendants != 272 {
		fmt.Printf("Expected story.Descendants to equal '%s' but was '%s'\n", "272", story.Descendants)
		t.Fail()
	}
	if story.Id != 13561388 {
		fmt.Printf("Expected story.Id to equal '%s' but was '%s'\n", "13561388", story.Id)
		t.Fail()
	}
	if len(story.Kids) != 2 {
		fmt.Printf("Expected story.Kids to have length of '2' but was'%s'\n", len(story.Kids))
		t.Fail()
	}
	if (story.Kids[0] != 13561709 || story.Kids[1] != 13561527) {
		fmt.Println("story.Kids did not have expected values!")
		t.Fail()
	}
	if story.Score != 324 {
		fmt.Printf("Expected story.Score to equal '%s' but was '%s'\n", "324", story.Score)
		t.Fail()
	}
	if story.Time != 1486142414 {
		fmt.Printf("Expected story.Time to equal '%s' but was '%s'\n", "1486142414", story.Time)
		t.Fail()
	}
	if story.Title != "Amazon soars to more than 341K employees, adding 110K people in a single year" {
		fmt.Printf("Expected story.Title to equal '%s' but was '%s'\n", "Amazon soars to more than 341K employees, adding 110K people in a single year",
			story.Title)
		t.Fail()
	}
	if story.Type != "story" {
		fmt.Printf("Expected story.Type to equal '%s' but was '%s'\n", "story", story.Type)
		t.Fail()
	}
	if story.Url != "http://www.geekwire.com/2017/amazon-soars-340k-employees-adding-110k-people-single-year/" {
		fmt.Printf("Expected story.Url to equal '%s' but was '%s'\n", "http://www.geekwire.com/2017/amazon-soars-340k-employees-adding-110k-people-single-year/",
			story.Url)
		t.Fail()
	}

}