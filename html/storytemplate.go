package html

import (
	"html/template"
	"io"

	"github.com/tjheslin1/YAHNCrawler/crawler"
)

type StoriesContext struct {
	Title   string
	Stories []*crawler.Story
}

// GenerateHTML TODO
func GenerateHTML(context StoriesContext, out io.Writer) {
	template, err := template.New("index").Parse(storyHTMLTemplate)
	check(err)

	err = template.Execute(out, context)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const storyHTMLTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
    {{range .Stories}}<div>{{ .By }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
</body>
</html>`
