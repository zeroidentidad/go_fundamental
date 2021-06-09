package main

import (
	"net/http"

	"github.com/temoto/robotstxt"
)

func main() {
	// Get the contents of robots.txt from packtpub.com
	resp, err := http.Get("https://www.packtpub.com/robots.txt")
	if err != nil {
		panic(err)
	}

	// Process the response using temoto/robotstxt
	data, err := robotstxt.FromResponse(resp)
	if err != nil {
		panic(err)
	}

	// Look for the definition in the robots.txt file that matches the default Go User-Agent string
	grp := data.FindGroup("Go-http-client/1.1")
	if grp != nil {
		testUrls := []string{
			"/all",
			"/all?search=Go",
			"/bundles",
			"/contact/",
			"/search/",
			"/user/password/",
		}

		for _, url := range testUrls {
			print("checking " + url + "...")

			// Test the path against the User-Agent group
			if grp.Test(url) == true {
				// These paths are all permissable
				println("OK")
			} else {
				// These paths are not
				println("X")
			}
		}
	}
}
