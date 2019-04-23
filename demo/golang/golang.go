package main

import (
	"github.com/sirupsen/logrus"
)

func testCap() {
	// 为何 slice 可以 cap, map 为何不可以？
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

func testDefer() {
	var i = 1
	defer logrus.Infoln("i1=", i)
	i++
	defer func(i int) {
		logrus.Infoln("i=", i)
	}(i)
}

func testDefer2() {
	// return 语句不是原子操作， 1. rval = xxx ,  2. ret，在这两步之间执行 defer
	// defer 的底层由实现 deferproc 、deferreturn

	f := func() int {
		i := 5
		defer func() {
			i++
		}()
		return i
	}()

	f1 := func() (result int) {
		defer func() {
			result++
		}()
		return 0
	}()

	f2 := func() (result int) {
		t := 5
		defer func() {
			t = t + 5
		}()
		return t
	}()

	f3 := func() (r int) {
		defer func() {
			r = r + 5
		}()
		return 1
	}()

	logrus.Infoln(f, f1, f2, f3)

}

func main() {
	// errNil()
	// testCap()
	// testDefer()
	testDefer2()
}
