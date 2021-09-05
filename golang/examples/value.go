package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}
func (t *T) M2() {
	t.Name = "name2"
}

func structtest() {
	t1 := T{"t1"}
	fmt.Println("M1调用前：", t1.Name)
	t1.M1()
	fmt.Println("M1调用后：", t1.Name)
	fmt.Println("M2调用前：", t1.Name)
	t1.M2()
	fmt.Println("M2调用后：", t1.Name)
}
func structtest2() {
	t2 := &T{"t2"}
	fmt.Println("M1调用前：", t2.Name)
	t2.M1()
	fmt.Println("M1调用后：", t2.Name)
	fmt.Println("M2调用前：", t2.Name)
	t2.M2()
	fmt.Println("M2调用后：", t2.Name)
}
func value() {
	structtest()
	structtest2()
}

func assertion() {
	var a interface{} = float32(3)
	b, ok := a.(float32)
	if ok {
		logrus.Infoln(b)
	} else {
		logrus.Infoln("error", ok)
	}
}
