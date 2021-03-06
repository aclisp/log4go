package main

import (
	"time"
)

import l4g "github.com/aclisp/log4go"

func main() {
	log := l4g.NewLogger()
	log.AddFilter("network", l4g.FINEST, l4g.NewSocketLogWriter("tcp", "127.0.0.1:12124"))

	// Run `nc -u -l -p 12124` or similar before you run this to see the following message
	for {
		log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
		time.Sleep(1 * time.Second)
	}

	// This makes sure the output stream buffer is written
	log.Close()
	select {}
}
