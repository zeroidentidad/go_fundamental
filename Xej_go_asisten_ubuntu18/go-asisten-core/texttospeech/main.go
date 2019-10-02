package texttospeech

import (
	"strconv"
	"strings"
)

// NewClient texttospeech
func NewClient(lang string) *Client {
	return &Client{
		Lang: lang,
	}
}

// WordsParseToURL url parse
func WordsParseToURL(words string, lang string) string {
	words = strings.Replace(words, " ", "%20", 10)
	return "https://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=" + strconv.Itoa(len(words)) + "&client=tw-ob&q=" + words + "&tl=" + lang
}
