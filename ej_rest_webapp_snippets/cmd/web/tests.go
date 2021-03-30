package web

import (
	"net/http"
	"time"
)

func TestHumanDate(tm time.Time) string {
	hd := humanDate(tm)
	return hd
}

func TestSecureHeaders(next http.Handler) http.Handler {
	return secureHeaders(next)
}
