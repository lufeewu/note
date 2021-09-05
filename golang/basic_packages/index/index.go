package main

import (
	"fmt"
	"index/suffixarray"
)

func testIndex() {
	index := suffixarray.New([]byte("banananana"))
	offsets := index.Lookup([]byte("ana"), 3)
	for k, v := range offsets {
		fmt.Println(k, v)
	}
}
func main() {
	testIndex()
}
