package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"./config"
	"./webdemo"
	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func reloadConfig(reloaders []func(*config.Config) error, filename string) error {

	// mutex

	conf, err := config.LoadFile(filename)
	if err != nil {
		return err
	}
	failed := false
	for _, rl := range reloaders {
		if err := rl(conf); err != nil {
			logrus.Errorf("failed to apply configuration, error: %v", err)
			failed = true
		}
	}
	if failed {
		return fmt.Errorf("one or more errors occurred while applying the new configuration (--config.file=%q)", filename)
	}
	return nil
}

func main() {

	cfg := &struct {
		ListenPort int
		ListenHost string
		ConfigPath string
		Debug      bool
	}{}

	app := kingpin.New("chat", "A command-line chat application.")
	app.Flag("debug", "Enable debug mode").BoolVar(&cfg.Debug)
	app.Flag("ListenHost", "listen host").Default("0.0.0.0").StringVar(&cfg.ListenHost)
	app.Flag("ListenPort", "listen port").Default("19090").IntVar(&cfg.ListenPort)
	app.Flag("ConfigPath", "config path").Default("./config_demo.yaml").StringVar(&cfg.ConfigPath)
	// kingpin.MustParse(app.Parse(os.Args[1:]))
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if cfg.Debug {
		logrus.Warnln("debug mode on")
	}

	type closeOnce struct {
		C     chan struct{}
		once  sync.Once
		Close func()
	}
	var g run.Group
	reloadReady := &closeOnce{
		C: make(chan struct{}),
	}
	reloadReady.Close = func() {
		reloadReady.once.Do(func() {
			close(reloadReady.C)
		})
	}

	webHandler := webdemo.New(cfg.Debug)
	webHandler.ListenAddress = fmt.Sprintf("%s:%d", cfg.ListenHost, cfg.ListenPort)

	reloaders := []func(conf *config.Config) error{
		webHandler.ApplyConfig,
	}

	{
		// initial configuration
		cancel := make(chan struct{})
		g.Add(
			func() error {
				if err := reloadConfig(reloaders, cfg.ConfigPath); err != nil {
					logrus.Errorf("failed to apply config to application, error: %v", err)
					reloadReady.Close()
					return err
				}
				logrus.Infoln("init end...")
				reloadReady.Close()
				<-cancel
				return nil
			},
			func(err error) {
				close(cancel)
			},
		)
	}

	{
		// web run
		ctxWeb, cancelWeb := context.WithCancel(context.Background())

		g.Add(func() error {
			if err := webHandler.Run(ctxWeb); err != nil {
				return err
			}
			return nil
		}, func(err error) {
			logrus.Infoln("web end...")
			cancelWeb()
		})
	}

	{
		// reload config
		hup := make(chan os.Signal, 1)
		signal.Notify(hup, syscall.SIGHUP)
		cancel := make(chan struct{})
		g.Add(func() error {
			for {
				select {
				case <-hup:
					if err := reloadConfig(reloaders, cfg.ConfigPath); err != nil {
						logrus.Errorf("failed to apply config to application, error: %v", err)
					}
				case rc := <-webHandler.Reload():
					if err := reloadConfig(reloaders, cfg.ConfigPath); err != nil {
						rc <- err
					}
					rc <- nil
				case <-cancel:
					return nil
				}
			}
		}, func(err error) {
			close(cancel)

		})
	}

	{
		// gracefull exit
		term := make(chan os.Signal, 1)
		signal.Notify(term, os.Interrupt, syscall.SIGTERM)
		cancel := make(chan struct{})
		g.Add(func() error {
			select {
			case <-term:
				logrus.Infoln("receive SIGTERM, exiting gracefully....")
				reloadReady.Close()
			case <-cancel:
				reloadReady.Close()
			}
			return nil
		}, func(err error) {
			close(cancel)
		})
	}

	if err := g.Run(); err != nil {
		logrus.Errorf("failed to run demo, error: %v", err)
		os.Exit(1)
	}
	logrus.Infoln("see you next time!")

}
