package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func nextTimer(hour int) *time.Timer {
	var res *time.Timer
	now := time.Now()
	if now.Hour() < hour {
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
		res = time.NewTimer(next.Sub(now))
	} else {
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), hour, 0, 0, 0, next.Location())
		res = time.NewTimer(next.Sub(now))
	}
	return res
}

func nextTimerSec(sec int) *time.Timer {
	var res *time.Timer
	now := time.Now()
	if now.Second() < sec {
		next := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), sec, 0, now.Location())
		res = time.NewTimer(next.Sub(now))
	} else {
		next := now.Add(time.Second * 60)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), sec, 0, next.Location())
		res = time.NewTimer(next.Sub(now))
	}
	return res
}

func timer() {
	var t16, t18, t20 *time.Timer
	// t16 = nextTimer(16)
	// t18 = nextTimer(18)
	// t20 = nextTimer(20)
	logrus.Infoln(time.Now())
	t16 = nextTimerSec(16)
	t18 = nextTimerSec(18)
	t20 = nextTimerSec(20)
	for {
		select {
		case <-t16.C:
			logrus.Infoln(time.Now(), "sec:16 timer")
			t16 = nextTimerSec(16)
		case <-t18.C:
			logrus.Infoln(time.Now(), "sec:18 timer")
			t18 = nextTimerSec(18)
		case <-t20.C:
			logrus.Infoln(time.Now(), "sec:20 timer")
			t20 = nextTimerSec(20)
		}
	}
}

func main() {
	timer()
}

