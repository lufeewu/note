package main

import (
	"expvar"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var visits = expvar.NewInt("visits")

func handler(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln(visits.Value())
	visits.Add(1)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/visit", handler)
	http.ListenAndServe(":18180", nil)
}
