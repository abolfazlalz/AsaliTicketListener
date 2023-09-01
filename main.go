package main

import (
	"bus_listener/bus"
	"fmt"
)

func main() {
	alibaba := bus.NewAlibaba()
	helper := bus.NewHelper(alibaba)

	quit := make(chan struct{})
	go helper.CheckInterval(alibaba.Tehran(), alibaba.Shahroud(), quit)
	go helper.CheckInterval(alibaba.Shahroud(), alibaba.Tehran(), quit)

	_, _ = fmt.Scanln()
}
