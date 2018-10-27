package main

import (
	"fmt"
	"log/syslog"
	"os"
	"time"
)

const socketDefault = "/var/run/rsyslog/dev/log"

func main() {
	socket := os.Getenv("RSYSLOG_SOCKET")
	if socket == "" {
		socket = socketDefault
	}

	logger, err := syslog.Dial("unixgram", socket, syslog.LOG_DEBUG, "rsyslog-app-go")
	if err != nil {
		panic(fmt.Sprintf("failed to create syslog writer: %s", err))
	}

	// host, err := os.Hostname()
	// if err != nil {
	// host = ""
	// }

	for {
		// now := time.Now()
		// msg := fmt.Sprintf("sample syslog app on %s at %d-%d-%d %d:%d:%d", host, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		// logger.Info(msg)
		logger.Info("some go msg")
		time.Sleep(time.Second * 5)
	}
}
