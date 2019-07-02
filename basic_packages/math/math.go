package main

import (
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"math/cmplx"
	"unsafe"
)

func testAbs() {
	fmt.Println(math.Abs(33))
}
func testBig() {
	a := big.NewInt(128)
	a.SetBytes([]byte{255, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 01})
	fmt.Println(a.String(), a.BitLen())

	b := big.NewRat(23, 44)
	fmt.Println(b.FloatString(33), b, b.RatString())
}
func testBits() {
	fmt.Println(unsafe.Sizeof(int64(3)), bits.Reverse(0x8000000000000000))
}

func testCmplx() {
	c := -2 + 2i
	fmt.Println(c)
	fmt.Println(cmplx.Abs(c - 2i))
	fmt.Println(cmplx.Abs(c))
}

func main() {
	testCmplx()
}
