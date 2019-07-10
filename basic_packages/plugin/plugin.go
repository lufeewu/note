package main

import (
	"fmt"
	"plugin"
)

func examplePlugin() {
	p, _ := plugin.Open("./aplugin/aplugin.so")
	add, _ := p.Lookup("Add")
	sub, _ := p.Lookup("Subtract")
	sum := add.(func(int, int) int)(11, 2)
	fmt.Println(sum)
	sum = sub.(func(int, int) int)(32, 143214)
	fmt.Println(sum)
}

func main() {
	examplePlugin()
}
