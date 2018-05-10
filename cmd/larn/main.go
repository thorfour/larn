package main

import (
	"flag"
	"fmt"
	"runtime/debug"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game"
)

func init() {
	flag.Lookup("stderrthreshold").Value.Set("FATAL") // Only log fatal logs to stderr
	flag.Parse()
}

func main() {
	defer flushLogs() // To ensure logs are flushed
	if err := game.New().Start(); err != nil {
		glog.Fatalf("game exited with error: %v", err)
	}
}

func flushLogs() {
	if r := recover(); r != nil {
		fmt.Println("Larn encountered an error")
		glog.Error(string(debug.Stack())) // Log the stack trace
	}
	glog.Flush()
}
