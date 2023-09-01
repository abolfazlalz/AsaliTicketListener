package main

import (
	"bus_listener/bus"
	"bus_listener/env"
	"bus_listener/notification"
	"fmt"
	"log"
)

func main() {
	if err := env.LoadEnvironment(); err != nil {
		log.Fatalf("Error during load enviroments: %v", err)
	}

	alibaba := bus.NewAlibaba()
	notif := notification.NewNotification()
	helper := bus.NewHelper(alibaba, notif)

	quit := make(chan struct{})
	go helper.CheckInterval(alibaba.Tehran(), alibaba.Shahroud(), quit)
	go helper.CheckInterval(alibaba.Shahroud(), alibaba.Tehran(), quit)

	go func(notif notification.Interface) {
		_ = notif.Send(notification.Message{Text: "System started"})
	}(notif)

	_, _ = fmt.Scanln()
}
