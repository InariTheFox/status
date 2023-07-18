package main

import (
	"github.com/inarithefox/status/cmd"
)

var stopped chan bool

func main() {
	stopped = make(chan bool, 1)
	cmd.Execute()
	<-stopped
}
