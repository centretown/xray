package main

import (
	"xray/tools"
)

func main() {
	control := make(chan int)
	runner := tools.NewRunner()
	runner.Run(control)
	// time.Sleep(time.Second * 10)
	// control <- 1
}
