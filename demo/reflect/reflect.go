package main

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

func testReflect() {
	var num float32 = 3.2321
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	logrus.Infof("type of pointer: %v, settability of pointer: %v", pointer.Type(), newValue.CanSet())

	if newValue.CanSet() {
		newValue.SetFloat(2222.23)
		logrus.Infof("new value of pointer: %v %v", num, newValue)
	}

}

func main() {
	testReflect()

}
