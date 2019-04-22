package main

import "github.com/sirupsen/logrus"

func testCap() {
	// 为何 slice 可以 cap，map 为何不可以？
	// map 底层实现是什么？
	var s []int
	logrus.Infoln(cap(s))
	s = append(s, 1)
	logrus.Infoln(cap(s))
}

func errNil() {
	// nil 只用在 pointer、channel、func、interface、map 或者 slice 类型中
	// error 底层实现是什么？
	var err error
	logrus.Infoln(err == nil)
	logrus.Infoln(err != nil)

	var trap interface{}

	logrus.Infoln(trap == nil)
	logrus.Infoln(trap != nil)

}

func main() {
	// errNil()
	testCap()
}
