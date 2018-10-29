package main

import (
	"fmt"
	"log/syslog"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	hook, err := lSyslog.NewSyslogHook("", "", syslog.LOG_INFO, "rsyslog-app-go")
	if err != nil {
		panic(err)
	}
	log.Hooks.Add(hook)

	host, err := os.Hostname()
	if err != nil {
		host = ""
	}

	for {
		now := time.Now()
		msg := fmt.Sprintf("sample syslog app on %s at %d-%d-%d %d:%d:%d", host, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		// logger.Info(msg)
		log.Info(msg)

		time.Sleep(time.Second * 3)
	}
}
