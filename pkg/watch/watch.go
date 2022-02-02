package watch

import (
	"fmt"
	"time"
	"watcher/pkg/notification"
	"watcher/pkg/server"
)

func Watch(server server_watcher.ServerWatcher, timeout int) {
	var err error
	for {
		err = server.CheckConnection()

		if server.IsWorking() {
			if err != nil {
				serverWentDown(server, err)
			}
			time.Sleep(time.Duration(timeout) * time.Second)
		} else {
			if err == nil {
				serverStartedUp(server)
			} else {
				server.IncrementOffTime()
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func serverWentDown(server server_watcher.ServerWatcher, err error) {
	server.SetWorking(false)
	fmt.Println("ServerWatcher " + server.GetName() + " not working")
	notification.SendErrorNotification(server.GetName(), server.GetAddr(), server.GetMentionsText(), err)
}

func serverStartedUp(server server_watcher.ServerWatcher) {
	server.SetWorking(true)
	fmt.Println("ServerWatcher " + server.GetName() + " is working again")
	notification.SendOkNotification(server.GetName(), server.GetAddr())
}
