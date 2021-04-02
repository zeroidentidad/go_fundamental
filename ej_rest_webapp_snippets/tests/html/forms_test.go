package html_test

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

// go test -run TestSignupUser -v
func TestSignupUser(t *testing.T) {
	// new application struct y test server end-to-end
	app := newTestApplication(t)
	ts := newTestServer(t, app.Routes())
	defer ts.Close()

	// GET request y extract CSRF
	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantBody     []byte
	}{
		{"Valid submission", "test", "test@example.com", "validPa$$word", csrfToken, http.StatusSeeOther, nil},
		{"Empty name", "", "test@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("Este campo no puede estar vacío")},
		{"Empty email", "test", "", "validPa$$word", csrfToken, http.StatusOK, []byte("Este campo no puede estar vacío")},
		{"Empty password", "test", "test@example.com", "", csrfToken, http.StatusOK, []byte("Este campo no puede estar vacío")},
		{"Invalid email (incomplete domain)", "test", "test@example.", "validPa$$word", csrfToken, http.StatusOK, []byte("Este campo no es válido")},
		{"Invalid email (missing @)", "test", "testexample.com", "validPa$$word", csrfToken, http.StatusOK, []byte("Este campo no es válido")},
		{"Invalid email (missing local part)", "test", "@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("Este campo no es válido")},
		{"Short password", "test", "test@example.com", "pa$$word", csrfToken, http.StatusOK, []byte("Este campo es demasiado corto (mínimo 10 caracteres)")},
		{"Duplicate email", "test", "dupe@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("La dirección ya está en uso")},
		{"Invalid CSRF Token", "", "", "", "wrongToken", http.StatusBadRequest, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			code, _, body := ts.post(t, "/user/signup", form)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body %s contain %q", body, tt.wantBody)
			}
		})
	}
}
