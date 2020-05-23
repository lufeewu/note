package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	var list = []int{3, 4, 5, 1, 32, 3, 5, 8, 10, 32}
	SelectSort(list)
	fmt.Println(list)
}
