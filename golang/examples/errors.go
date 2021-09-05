package main

import "strconv"

type errors struct {
	a int64
}

func (e *errors) Error() string {
	return "test" + strconv.Itoa(int(e.a))
}
