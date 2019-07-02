package main

import (
	"fmt"
	"sort"
)

type person struct {
	ID   int
	Age  int
	Name string
}

type sortByAge []person

func (a sortByAge) Len() int           { return len(a) }
func (a sortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func sortAge() {
	a := sortByAge{
		{0, 2, "Bob"},
		{1, 21, "Alice"},
		{2, 16, "Jack"},
	}

	// default
	fmt.Println("before sort:", a)

	// sort
	sort.Sort(a)
	fmt.Println("sort:", a)

	// reverse
	sort.Sort(sort.Reverse(a))
	fmt.Println("reverse sort:", a)

}
func main() {
	sortAge()

}
