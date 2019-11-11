package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func terrors() {
	var err error
	err = &errors{
		73,
	}
	logrus.Infoln(err)
}

func tsocket() {
	cfg := &struct {
		Type    string
		Address string
	}{}

	app := kingpin.New("mcenter", "A docker containers manager tool.")
	app.Flag("type", "socket type").Default("client").StringVar(&cfg.Type)
	app.Flag("address", "listen address").Default(":10208").StringVar(&cfg.Address)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	switch cfg.Type {
	case "server":
		logrus.Infoln("start socket server...")
		socketServer(cfg.Address)
	case "client":
		logrus.Infoln("start socket client...")
		socketClient(cfg.Address)
	case "clientBench":
		logrus.Infoln("start socket many client...")
		for i := 0; i < 111201; i++ {
			go socketClient(cfg.Address)
			time.Sleep(300 * time.Millisecond)
		}
		socketClient(cfg.Address)
	default:
		logrus.Warnln("wrong socket type, exit...")
	}
}

func testSlice() {
	var result = make([]int, 3, 4)
	var result2 = make([]int, 3)
	fmt.Println(result, len(result), cap(result))
	fmt.Println(append(result, 1))
	fmt.Println(result2, len(result2), cap(result2))
	_ = append(result, 1)
	result2 = append(result, 1, 2)
	fmt.Println(result, len(result), cap(result))
	fmt.Println(result2, len(result2), cap(result2))

	fmt.Println(append(result, 1))
	_ = append(result, 1)
	// result2 = append(result, 1, 2)
	// append(result, 1)
	result2 = append(result, 1)
	fmt.Println(result2, result)
	result[0] = 1
	fmt.Println(result2, result)

}

func main() {
	testSlice()
}
