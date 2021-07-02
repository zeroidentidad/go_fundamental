package main

import (
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
)

func main() {
	doc, err := htmlquery.LoadURL("https://gophers-latam.github.io/posts/")
	if err != nil {
		panic(err)
	}

	dealTextNodes := htmlquery.Find(doc, `//div[@class="post-list"]//text()`)

	if err != nil {
		panic(err)
	}

	println("Here is posts of the day!")
	println("----------------------------------")

	for _, node := range dealTextNodes {
		text := strings.TrimSpace(node.Data)
		matchTagNames, _ := regexp.Compile("^(div|article|h3)$")
		text = matchTagNames.ReplaceAllString(text, "")
		if text != "" {
			println(text)
		}
	}
}
