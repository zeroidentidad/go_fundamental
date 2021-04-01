package mock_test

import (
	"bytes"
	"net/http"
	"testing"
)

// go test -run TestShowSnippet -v
func TestShowSnippet(t *testing.T) {
	// new application struct con mocked deps
	app := newTestApplication(t)

	// new test server end-to-end
	ts := newTestServer(t, app.Routes())
	defer ts.Close()

	// table-driven tests check responses diferentes URLs
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("Un viejo estanque silencioso...")},
		{"Non-existent ID", "/snippet/3", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash ID", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body %q", tt.wantBody)
			}
		})
	}

}
