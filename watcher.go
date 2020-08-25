package main

import (
	"fmt"
	"time"
)

func watch(server Server, timeout int) {
	var err error
	for {
		err = server.checkConnection()
		if server.isWorking() {
			if err != nil {
				server.setWorking(false)
				fmt.Println("Server " + server.getName() + " not working")
				sendErrorNotification(server.getName(), server.getAddr(), server.getMentionsText(), err)
			}
			time.Sleep(time.Duration(timeout) * time.Second)
		} else {
			if err == nil {
				server.setWorking(true)
				fmt.Println("Server " + server.getName() + " is working again")
				sendOkNotification(server.getName(), server.getAddr())
			} else {
				server.incrementOffTime()
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
