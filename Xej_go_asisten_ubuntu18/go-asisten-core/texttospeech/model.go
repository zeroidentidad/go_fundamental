package texttospeech

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Lang string
}

func (client *Client) ConvertToFile(texto string, filename string) error {
	resp, err := client.Convert(texto)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, resp, 0644)
}

func (client *Client) ConvertToBase64(texto string) (string, error) {
	resp, err := client.Convert(texto)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(resp), nil
}

func (client *Client) Convert(texto string) ([]byte, error) {
	resp, err := http.Get(WordsParseToURL(texto, client.Lang))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
