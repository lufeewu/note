package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func testIO() {
	f, err := os.Open("./tmp")
	if err != nil {
		logrus.Error(err)
		return
	}
	f.Seek(100, 0)
	buffer := make([]byte, 1024)
	n, err := f.Read(buffer)
	if err != nil {
		logrus.Errorf("read error %v", err)
		return
	}
	curOffset, _ := f.Seek(0, os.SEEK_CUR)
	logrus.Infof("read buffer len= %d, cur_offset is %d", n, curOffset)
	logrus.Infoln(string(buffer))
}

type readerInterface interface {
	Read(p []byte) (int64, error)
}

type readerStruct struct {
	readerInterface
}

func (r readerStruct) Read(p []byte) (int64, error) {
	return 0, nil
}

type impStruct struct {
	r readerInterface
	v int64
}

func testInterface() {
	// 1. 实现 interface
	/*
		var r readerStruct
		var s = impStruct{
			r: r,
			v: 8,
		}
		var b []byte
		logrus.Infoln(s.v)
		logrus.Infoln(s.r.Read(b))
	*/

	// 2. 实现 interface、值方法、指针方法
	// Go 语言实现一个接口并不需要显示声明，只要你实现了接口中的所有方法就认为你实现了这个接口 Duck typing
	// 指针方法集合包括值方法接口以及指针方法，值类型方法集合是指针类型方法的真子集

	var sr = &io.SectionReader{}
	var b io.Reader
	b = sr
	var p []byte
	logrus.Infoln(b.Read(p))
	logrus.Infoln(string(p))

}

/*
	实现了 io 的类型
	bufio、bytes、strings、crypto/tls、archive/tar、math/rand、os/File
*/

func main() {
	// testIO()
	testInterface()
}
