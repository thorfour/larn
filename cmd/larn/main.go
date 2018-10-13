package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
	"github.com/thorfour/larn/pkg/game"
	"github.com/thorfour/larn/pkg/game/data"
)

var (
	difficulty = flag.Int("d", 0, "sets the game difficulty")
)

func init() {
	flag.Parse()
	logfile, err := ioutil.TempFile("", "larn.log")
	if err != nil {
		log.Errorf("unable to open log file: %v", err)
		return
	}

	log.SetOutput(logfile)
}

func main() {
	defer flushLogs() // To ensure logs are flushed
	if err := game.New(&data.Settings{
		Difficulty: *difficulty,
	}).Start(); err != nil {
		log.WithField("error", err).Fatal("game exited with error")
	}
}

func flushLogs() {
	if r := recover(); r != nil {
		fmt.Println("Larn encountered an error")
		log.Error(string(debug.Stack())) // Log the stack trace
	}
}
