package mock_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"pastein/cmd/web"
	"pastein/pkg/mysql/mock"
	"testing"
	"time"

	"github.com/golangcollege/sessions"
)

type testServer struct {
	*httptest.Server
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
