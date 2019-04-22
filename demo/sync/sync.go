package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

func syncPool() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	a1 := p.Get().(int)
	a2 := p.Get().(int)
	p.Put(3)
	b := p.Get().(int)
	logrus.Infoln(a, a1, a2, b)
}

func syncWaitGroup() {
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			logrus.Infoln(i)
			w.Done()
		}(i)
	}

	w.Wait()

}
func main() {
	syncWaitGroup()

}
