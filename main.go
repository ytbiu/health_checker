package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"health_checker/config"
	"health_checker/router"
	"health_checker/service/common"
	"time"
)

var mode string

func init() {
	flag.StringVar(&mode, "mode", "debug", "mode")
}

func main() {
	flag.Parse()
	logrus.Info("mode : ", mode)
	if mode != gin.DebugMode && mode != gin.TestMode && mode != gin.ReleaseMode {
		panic(fmt.Sprintf("invalid mode : %s", mode))
	}

	config.Init(mode)

	//gin.SetMode(mode)
	r := gin.Default()
	router.Init(r)

	go func() {
		periodSecond := time.Duration(config.ConfigInfo.NvidiaFetchPeriodSeconds) * time.Second
		logrus.Infof("fetchNvidia will exec every %s", periodSecond)
		for {
			common.FetchNvidia()
			time.Sleep(periodSecond)
		}
	}()

	r.Run(config.ConfigInfo.ListenAddr)
}
