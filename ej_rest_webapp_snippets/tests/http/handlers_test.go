package http_test

import (
	"net/http"
	"testing"
)

// go test -run TestPing -v
func TestPing(t *testing.T) { // End-To-End Test
	// new app test server
	app := newTestApplication(t)
	ts := newTestServer(t, app.Routes())
	defer ts.Close()

	// request test a test_server
	code, _, body := ts.get(t, "/ping")

	// check status code
	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	// check response body
	if string(body) != "OK" {
		t.Errorf("want body %q", "OK")
	}
}
