package http_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"pastein/cmd/web"
	"testing"
)

// go test -run TestSecureHeaders -v
func TestSecureHeaders(t *testing.T) {
	// new httptest.ResponseRecorder y http.Request
	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// mock HTTP handler para pasar a secureHeaders
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	web.TestSecureHeaders(next).ServeHTTP(rr, r)

	// result method para obtener http.Response
	rs := rr.Result()

	// check middleware set the X-Frame-Options header
	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("want %q; got %q", "deny", frameOptions)
	}

	// check middleware set X-XSS-Protection header
	xssProtection := rs.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("want %q; got %q", "1; mode=block", xssProtection)
	}

	// check middleware llama next handler in line
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	// check response body
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body %q", "OK")
	}

}
