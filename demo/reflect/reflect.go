package main

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

func testReflect() {
	var num float32 = 3.2321
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	// logrus.Infof("type of pointer: %v, settability of pointer: %v", pointer.Type(), newValue.CanSet())

	if newValue.CanSet() {
		newValue.SetFloat(2222.23)
		// logrus.Infof("new value of pointer: %v %v", num, newValue)
	}
}

type at struct {
	Am string `json:"name"`
	By int    `json:"by"`
	T3 bool
}

// reflect 不能获取私有变量的 value
func testReflect2() {
	var a *at
	a = new(at)
	objType := reflect.TypeOf(*a)
	objValue := reflect.ValueOf(*a)
	for i := 0; i < objType.NumField(); i++ {
		logrus.Infoln(objType.Field(i).Name, objType.Field(i).Type, objValue.Field(i).Interface())
		field := objType.Field(i).Tag.Get("json")
		logrus.Infoln("json:", field, len(field))
	}
	logrus.Infoln(objValue.Interface())

	var err interface{} = nil
	var nil2 interface{} = nil
	var nil3 *int = nil
	logrus.Infoln(err == nil2, nil2 == nil3)

}

func main() {
	testReflect()
	testReflect2()
}
