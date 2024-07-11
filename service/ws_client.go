package service

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"health_checker/config"
	"net/url"
	"sync"
)

var reportToClusterWsConn *websocket.Conn

func DialWsServer() {
	u := url.URL{Scheme: "ws", Host: config.ConfigInfo.DBCHealthCheckClusterAddr, Path: config.ConfigInfo.DBCHealthCheckReportPath}
	logrus.Infof("Connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logrus.Info("dial:", err)
	}
	reportToClusterWsConn = c
}

var once sync.Once

func GetReportToClusterWsConn() *websocket.Conn {
	once.Do(func() {
		DialWsServer()
	})
	return reportToClusterWsConn
}
