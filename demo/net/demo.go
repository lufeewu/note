package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func httpServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		logrus.Infof("EscapedPath:%s, Path:%s, RawPath:%s",
			c.Request.URL.EscapedPath(), c.Request.URL.Path, c.Request.URL.RawPath)
		c.JSON(200, nil)
	})
	t := r.Group("/terra")
	t.GET("/ping", func(c *gin.Context) {
		logrus.Infof("EscapedPath:%s, Path:%s, RawPath:%s",
			c.Request.URL.EscapedPath(), c.Request.URL.Path, c.Request.URL.RawPath)
		c.JSON(200, nil)
	})

	logrus.Infoln("r")
	r.Run(":18080")
}

func main() {
	httpServer()
}
