package watch

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"watcher/pkg/notification"
	"watcher/pkg/serverwatcher"
)

func Watch(watcher serverwatcher.ServerWatcher) {
	timeout := watcher.GetTimeout()
	var err error
	for {
		err = watcher.CheckConnection()

		if watcher.IsWorking() {
			if err != nil {
				serverWentDown(watcher)
			}
			time.Sleep(time.Duration(timeout) * time.Second)
		} else {
			if err == nil {
				serverStartedUp(watcher)
			} else {
				watcher.IncrementOffTime()
				if shouldReportAgain(watcher) {
					notification.ServerStillIsDown(watcher)
				}
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func shouldReportAgain(server serverwatcher.ServerWatcher) bool {
	// todo эту уникальную формулу явно стоит переделать
	return server.GetOffTime()&0b111 == 0b111 && (server.GetOffTime() < 30 || rand.Intn(4) == 1)
}

func serverWentDown(watcher serverwatcher.ServerWatcher) {
	watcher.SetWorking(false)
	log.Info("Server " + watcher.GetFormattedName() + " not working")
	notification.ServerWentDown(watcher)
}

func serverStartedUp(watcher serverwatcher.ServerWatcher) {
	watcher.SetWorking(true)
	log.Info("Server " + watcher.GetFormattedName() + " is working again")
	notification.ServerIsUp(watcher)
}
