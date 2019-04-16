package main

import (
	"fmt"
)

type Birds interface {
	Twitter() string
	Fly(high int) bool
}

type Chicken interface {
	Birds
	Walk()
}

type Sparrow struct {
	name string
}

func (s *Sparrow) Fly(high int) bool {
	fmt.Printf("fly %d\n", high)
	return true
}

func (s *Sparrow) Twitter() string {
	return fmt.Sprintf("%s,jojojo", s.name)
}

func (s Sparrow) Walk() {
	// ...
}

func BirdAnimation(bird Birds, high int) {
	fmt.Printf("BirdAnimation of %T\n", bird)
	fmt.Println(bird.Twitter())
	bird.Fly(high)
}

func ChickenAnimation(chicken Chicken) {
	fmt.Printf("ChickenAnimation of %T\n", chicken)
	chicken.Walk()
	chicken.Twitter()
}

func checkassign() {
	// var bird Birds
	// sparrow := &Sparrow{name: "hey"}
	// bird = sparrow
	// BirdAnimation(bird, 1000)
	// BirdAnimation(sparrow, 1000)
	var chicken Chicken
	sparrow2 := Sparrow{}
	chicken = &sparrow2 // 一个指针类型的方法列表必然包含所有接收者为指针接收者的方法，同理非指针类型的方法列表也包含所有接收者为非指针类型的方法。在我们例子中*Sparrow首先包含：Fly和Twitter；Sparrow包含Walk
	ChickenAnimation(chicken)
}
