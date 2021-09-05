package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func testConvert() {
	type MyInt int

	a := []MyInt{0, 1, 2}
	// b := ([]int)(a) // error: cannot convert a (type []MyInt) to type []int
	b := *(*[]int)(unsafe.Pointer(&a))

	b[0] = 3

	fmt.Println("a =", a) // a = [3 1 2]
	fmt.Println("b =", b) // b = [3 1 2]

	a[2] = 9

	fmt.Println("a =", a) // a = [3 1 9]
	fmt.Println("b =", b) // b = [3 1 9]
}

var data *string

// get data atomically
func Data() string {
	p := (*string)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
	))
	if p == nil {
		return ""
	} else {
		return *p
	}
}

// set data atomically
func SetData(d string) {
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
		unsafe.Pointer(&d),
	)
}

func testSyncAtomic() {
	var wg sync.WaitGroup
	wg.Add(8)

	for range [4]struct{}{} {
		go func() {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)

			log.Println(Data())
			wg.Done()
		}()
	}

	for i := range [4]struct{}{} {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)
			s := fmt.Sprint("#", i)
			log.Println("====", s)

			SetData(s)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("final data = ", *data)
}

func main() {
	// testSyncAtomic()
}
