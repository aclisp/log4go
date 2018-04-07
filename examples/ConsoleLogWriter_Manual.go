package main

import (
	"time"
)

import l4g "github.com/aclisp/log4go"

func main() {
	log := l4g.NewLogger()
	log.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())
	for {
		log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
		time.Sleep(1 * time.Second)
	}
}
