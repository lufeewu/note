package main

import (
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

func main() {
	tsocket()
}
