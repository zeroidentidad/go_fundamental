package main

import "time"

func main() {

	a := make(chan string)
	b := make(chan string)
	c := make(chan string)
	d := make(chan string)

	in1 := "in1"
	in2 := "in2"
	in3 := "in3"
	in4 := "in4"

	go func(in string) {
		a <- testfunctionA(in)
	}(in1)

	go func(in string) {
		b <- testfunctionB(in)
	}(in2)

	go func(in string) {
		c <- testfunctionC(in)
	}(in3)

	go func(in string) {
		d <- testfunctionD(in)
	}(in4)

	finalterm := PrintAllfunction(<-a, <-b, <-c, <-d)
	println(finalterm)
}

func testfunctionA(input string) string {
	time.Sleep(9000)
	println(input)
	return input
}

func testfunctionB(input string) string {
	time.Sleep(6000)
	println(input)
	return input
}
func testfunctionC(input string) string {
	time.Sleep(3000)
	println(input)
	return input
}
func testfunctionD(input string) string {
	time.Sleep(1000)
	println(input)
	return input
}

func PrintAllfunction(a string, b string, c string, d string) string {
	println("-----------------")
	println(a)
	println(b)
	println(c)
	println(d)
	return "done"
}
