package main

import (
	"sync"
	"time"

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

func syncChan() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 100)
	for i := 0; i < 1203; i++ {
		ch <- 1
		wg.Add(1)
		go func(i int) {
			defer func() {
				<-ch
				logrus.Infoln("over", i)
				wg.Done()
			}()
			time.Sleep(1 * time.Second)
		}(i)
	}

	wg.Wait()
}

func main() {
	// syncWaitGroup()
	syncChan()

}
