package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	// "log/syslog"
	syslog "github.com/RackSec/srslog"
)

// const socketDefault = "/var/run/rsyslog/dev/log"

func getLogger() (*syslog.Writer, error) {
	for {
		logger, err := syslog.Dial("unixgram", "/dev/log", syslog.LOG_DEBUG, "rsyslog-app-go")
		if err != nil {
			fmt.Printf("failed to create syslog writer: %s", err)
			time.Sleep(time.Second * 5)
			continue
		}
		return logger, nil
	}
}

func main() {
	// socket := os.Getenv("RSYSLOG_SOCKET")
	// if socket == "" {
	// 	socket = socketDefault
	// }

	printDirContent("/dev")
	printDirContent("/var/run/rsyslog")

	logger, err := syslog.Dial("unixgram", "/dev/log", syslog.LOG_DEBUG, "rsyslog-app-go")
	if err != nil {
		panic(fmt.Sprintf("failed to create syslog writer: %s", err))
	}

	// logger.SetFormatter(syslog.DefaultFormatter)
	// logger.SetFormatter(syslog.RFC3164Formatter)
	logger.SetFormatter(syslog.RFC5424Formatter)

	// host, err := os.Hostname()
	// if err != nil {
	// host = ""
	// }

	for {
		// now := time.Now()
		// msg := fmt.Sprintf("sample syslog app on %s at %d-%d-%d %d:%d:%d", host, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		// logger.Info(msg)
		logger.Info("some go msg")
		time.Sleep(time.Second * 15)
	}
}

func printDirContent(dir string) {
	err := filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})

	if err != nil {
		fmt.Printf(" :: err: %v\n", err)
	}
}
