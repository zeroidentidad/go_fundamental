package web

import (
	"time"
)

func TestHumanDate(tm time.Time) string {
	hd := humanDate(tm)
	return hd
}
