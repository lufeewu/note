package main

import (
	"bytes"
	"fmt"
	"time"
	"unsafe"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

func selectCase() {
	ch := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		ch[i] = make(chan int)
	}
	go func() {
		i := 0
		for {
			logrus.Infof("i=%d", i)
			time.Sleep(1 * time.Second)
			ch[i] <- 1
			i = (i + 1) % 3
		}
	}()

	for {

		select {
		case <-ch[0]:
			logrus.Infoln("0")
		case <-ch[1]:
			logrus.Infoln("1")
		case <-ch[2]:
			logrus.Infoln("2")
		default:
			logrus.Infoln("default")
			time.Sleep(1 * time.Second)
		}
	}
}

func testSize() {
	// selectCase()
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	var a []int
	var b map[int]int
	var c struct{}
	fmt.Println(unsafe.Sizeof(a)) // 24
	fmt.Println(unsafe.Sizeof(b)) // 8
	fmt.Println(unsafe.Sizeof(c)) // 0
	var ac uintptr
	fmt.Println(unsafe.Sizeof(ac)) // 8
}

// Time chan test
func Time() {
	type test struct {
		Cancel chan struct{}
	}
	t := test{
		Cancel: make(chan struct{}),
	}
	go func() {
		<-t.Cancel
		logrus.Infoln("1")
	}()
	time.Sleep(1 * time.Second)
	t.Cancel <- struct{}{}
	time.Sleep(1 * time.Second)
}

// GO 语言核心36讲-39 bytes 包的泄漏
func unreadBytesTest() {
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("%q %d\n", contents, buffer1.Cap())
	unreadBytes := buffer1.Bytes()
	fmt.Printf("%v \n", unreadBytes)

	buffer1.WriteString("cdefg")
	fmt.Printf("The Capacity of buffer:%d %v\n", buffer1.Cap(), string(buffer1.Bytes()))
	unreadBytes = unreadBytes[:cap(unreadBytes)] // leak new string from unreadBytes
	fmt.Printf("%v\n", string(unreadBytes))
}

func letsEncrypt() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	logrus.Infoln(autotls.Run(r, "127.0.0.1"))

}

func testClose() {

	cancel := make(chan struct{})
	cancel2 := make(chan struct{})

	go func() {
		// close(cancel2)
		time.Sleep(1 * time.Second)
		cancel <- struct{}{}
	}()

	select {
	case _, ok := <-cancel:
		logrus.Infoln("cancel1", ok)
	case _, ok := <-cancel2:
		logrus.Infoln("cancel2", ok)
	}

}

func main() {
	// Time()
	// t := time.NewTimer(10 * time.Second)
	// for {
	// 	t.Reset(1 * time.Second)
	// 	logrus.Infoln("1")
	// 	select {
	// 	case <-t.C:
	// 		logrus.Infoln("2")
	// 	}
	// }
	// unreadBytesTest()

	// letsEncrypt()
	// router := gin.Default()

	// router.RunTLS(":18080", "", "")
	testClose()

}
