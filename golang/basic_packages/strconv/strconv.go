package main

import "fmt"

func testString() {
	var a string
	fmt.Println(a)
	a = "32"
	fmt.Println(a)
	// a[0] = a[0] + 1
	fmt.Println(string(a[0] + a[1]))

}
func main() {
	testString()
}
