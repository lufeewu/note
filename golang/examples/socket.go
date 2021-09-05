package main

import (
	"bufio"
	"net"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	conns map[string]net.Conn
)

func connServerHandler(c net.Conn) {
	if c == nil {
		return
	}
	conns[c.RemoteAddr().String()] = c
	defer delete(conns, c.RemoteAddr().String())
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}

		inStr := strings.TrimSpace(string(buf[0:cnt]))
		inputs := strings.Split(inStr, " ")
		switch inputs[0] {
		case "ping":
			logrus.Infof("ping from %s, response: %s", c.RemoteAddr(), c.LocalAddr())
			c.Write([]byte("ping " + c.LocalAddr().String()))
		case "echo":
			echoStr := strings.Join(inputs[1:], " ")
			logrus.Infof("%s: echo %s", c.RemoteAddr(), echoStr)
			response := []byte(echoStr)
			if len(response) == 0 {
				response = []byte("nil")
			}
			c.Write(response)
		case "quit":
			c.Close()
			break
		default:
			unsupport := strings.Join(inputs[:], " ")
			logrus.Warnf("unsupported command from %s: %s", c.RemoteAddr(), unsupport)
			c.Write([]byte("unsupport command: \"" + unsupport + "\""))
		}
	}
	logrus.Infof("connection from %v closed. ", c.RemoteAddr())
}

func socketServer(address string) {
	conns = make(map[string]net.Conn)
	server, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Infof("failed to start server, error: %v", err)
		return
	}
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		for {
			<-ticker.C
			logrus.Infof("broadcast to %d begin...", len(conns))
			for _, c := range conns {
				// logrus.Infof("send to %s", c.RemoteAddr())
				_, err := c.Write([]byte(time.Now().String()))
				if err != nil {
					logrus.Warnf("failed to send, error: %v", err)
				}
			}
		}
	}()

	logrus.Infoln("Server started ...")
	for {
		conn, err := server.Accept()
		if err != nil {
			logrus.Warnf("failed to connect, error: %v", err)
			break
		}
		go connServerHandler(conn)

	}
}

func connClientHandler(c net.Conn) error {
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)

	var err error
	// wait for message
	go func() error {
		for {
			var cnt int
			cnt, err = c.Read(buf)
			if err != nil {
				logrus.Warnf("failed to read data, error: %v", err)
				return err
			}
			logrus.Infoln(string(buf[0:cnt]))
		}
	}()
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			return err
		}
		c.Write([]byte(input))
	}
}
func socketClient(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		logrus.Warnf("failed to connect, error:%v", err)
		return err
	}
	err = connClientHandler(conn)
	return err
}
