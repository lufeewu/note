package webdemo

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"../config"
)

type WebDemo struct {
	ListenAddress string
	config        *config.Config
	router        *gin.Engine
	reloadChan    chan chan error
}

func New(debug bool) *WebDemo {

	if debug == false {
		gin.SetMode(gin.ReleaseMode)
	}
	w := &WebDemo{
		reloadChan: make(chan chan error),
		router:     gin.New(),
	}
	w.Register()
	return w
}

func (w *WebDemo) ApplyConfig(conf *config.Config) error {
	w.config = conf

	logrus.Infof("%p %p", w.config, conf)
	return nil
}

func (w *WebDemo) Register() {

	w.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 0, "message": "success ping"})
	})

	w.router.GET("/reloadconfig", func(c *gin.Context) {
		rc := make(chan error)
		w.reloadChan <- rc
		if err := <-rc; err != nil {
			c.JSON(200, gin.H{"code": 1001, "message": "reload config error"})
		} else {
			c.JSON(200, gin.H{"code": 0, "message": *w.config})
		}
	})
}

func (w *WebDemo) Run(ctx context.Context) error {

	srv := &http.Server{
		Addr:    w.ListenAddress,
		Handler: w.router,
	}

	errCh := make(chan error)
	go func() {
		errCh <- srv.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		if err := srv.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	}
}

func (w *WebDemo) Reload() <-chan chan error {
	return w.reloadChan
}
