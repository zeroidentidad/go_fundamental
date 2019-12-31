package main

func main() {
	testreturn1, testreturn2 := testfunc("test ", "test ")
	println(testreturn1 + testreturn2)
}

/*func testfunc(tst string, tst1 string) {
	println(tst + tst1)
}*/

/*func testfunc(tst string, tst1 string) string {
	return tst + tst1
}*/

func testfunc(tst string, tst1 string) (string, string) {
	return tst + "1 ", tst1 + "2"
}
