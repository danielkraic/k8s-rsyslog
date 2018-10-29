package main

import (
	"fmt"
	"os"
	"time"

	// syslog "github.com/RackSec/srslog"
	"log/syslog"

	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

// const socketDefault = "/var/run/rsyslog/dev/log"

// func getLogger() *syslog.Writer {
// 	for {
// 		logger, err := syslog.Dial("tcp", "rsyslog:514", syslog.LOG_DEBUG, "rsyslog-app-go")
// 		if err != nil {
// 			fmt.Printf("failed to create syslog writer: %s", err)
// 			time.Sleep(time.Second * 5)
// 			continue
// 		}
// 		return logger
// 	}
// }

func main() {
	// logger := getLogger()
	// logger.SetFormatter(syslog.RFC5424Formatter)

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	hook, err := lSyslog.NewSyslogHook("udp", "rsyslog:514", syslog.LOG_INFO, "rsyslog-app-go")

	if err == nil {
		log.Hooks.Add(hook)
	}

	host, err := os.Hostname()
	if err != nil {
		host = ""
	}

	for {
		now := time.Now()
		msg := fmt.Sprintf("sample syslog app on %s at %d-%d-%d %d:%d:%d", host, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		// logger.Info(msg)
		log.Info(msg)

		time.Sleep(time.Second * 15)
	}
}
