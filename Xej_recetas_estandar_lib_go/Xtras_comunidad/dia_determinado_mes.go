package main

import (
	"fmt"
	"time"
)

func main() {

	year, month, _ := time.Now().Date()
	afu := time.Date(year, month, LastDay(time.Now()), 0, 0, 0, 0, time.Local)
	fmt.Printf("%v\n", afu.Add(-time.Duration(((afu.Weekday()+7)-time.Monday)%7)*24*time.Hour))

	// tests
	date_test1 := time.Date(2020, time.September, 30-7 /*LastDay()*/, 0, 0, 0, 0, time.Local)
	date_test2 := time.Date(2020, time.August, 31-8, 0, 0, 0, 0, time.Local)
	date_test3 := time.Date(2020, time.February, 29 /*28*/ -7, 0, 0, 0, 0, time.Local)
	fmt.Printf("%v\n", date_test1.Add(-time.Duration(((date_test1.Weekday()+7)-time.Monday)%7)*24*time.Hour))
	fmt.Printf("%v\n", date_test2.Add(-time.Duration(((date_test2.Weekday()+7)-time.Monday)%7)*24*time.Hour))
	fmt.Printf("%v\n", date_test3.Add(-time.Duration(((date_test3.Weekday()+7)-time.Monday)%7)*24*time.Hour))
}

func LastDay(t time.Time) int {
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)
	day := lastDay.Day()
	if day == 31 {
		return day - 8
	}
	return day - 7
}
