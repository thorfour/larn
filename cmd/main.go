package main

import (
	"flag"

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
	// TODO panic greacefully if need be
	glog.Flush()
}
