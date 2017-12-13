package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game"
)

func init() {
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
	}
	glog.Flush()
}
