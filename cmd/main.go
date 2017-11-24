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
	if err := game.New().Start(); err != nil {
		glog.Fatalf("game exited with error: %v", err)
	}
}
