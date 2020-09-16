package async

// FetchAll toma una lista de URLs
func FetchAll(urls []string, c *Client) {
	for _, url := range urls {
		go c.AsyncGet(url)
	}
}
