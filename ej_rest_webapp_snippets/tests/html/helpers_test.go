package html_test

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"pastein/cmd/web"
	"pastein/pkg/mysql/mock"
	"regexp"
	"testing"
	"time"

	"github.com/golangcollege/sessions"
)

type testServer struct {
	*httptest.Server
}

// CSRF token en HTML
var csrfTokenRX = regexp.MustCompile(`'csrf_token' value='(.+)'`)

func extractCSRFToken(t *testing.T, body []byte) string {
	// extract the token del HTML body.
	matches := csrfTokenRX.FindSubmatch(body)
	if len(matches) < 1 {
		t.Fatal("no csrf token found in body")
	}

	// evitar transformación de caracteres a su formato de escape
	return html.UnescapeString(string(matches[1]))
}

func newTestApplication(t *testing.T) *web.Application {
	// instancia template cache
	templateCache, err := web.NewTemplateCache("./../../ui/html/")
	if err != nil {
		t.Fatal(err)
	}

	// instancia session manager como real
	session := sessions.New([]byte("z3Roh+pPbnzHbS*+9Pk8qGWhTzbpa@jf"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	return &web.Application{
		ErrorLog:      log.New(ioutil.Discard, "", 0),
		InfoLog:       log.New(ioutil.Discard, "", 0),
		Session:       session,
		Snippets:      &mock.SnippetModel{},
		TemplateCache: templateCache,
		Users:         &mock.UserModel{},
	}
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	// new cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	// set cookie jar al client
	ts.Client().Jar = jar

	// disable redirect-following
	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

// helper test GET routes
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

// helper test POST routes
func (ts *testServer) post(t *testing.T, urlPath string, form url.Values) (int, http.Header, []byte) {
	rs, err := ts.Client().PostForm(ts.URL+urlPath, form)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}
