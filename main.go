package main

import (
	"fmt"
	"os"
	"searcher/src/config"
	"searcher/src/core"
	"time"
)

func main() {
	go loading()

	var out string

	if len(os.Args) <= 1 {
		out = "null"
	} else {
		out = os.Args[1]
	}
	config := config.ReadConfig()
	core.Launch(config, out)
}

func loading() {
	fmt.Printf("Loading.\r")
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Loading..\r")
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Loading...\r")
	time.Sleep(300 * time.Millisecond)
}
