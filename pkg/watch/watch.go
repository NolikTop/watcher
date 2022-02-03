package watch

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"watcher/pkg/notification"
	"watcher/pkg/server"
)

func Watch(serv server.Server) {
	timeout := serv.GetTimeout()
	var err error
	for {
		err = serv.CheckConnection()

		if serv.IsWorking() {
			if err != nil {
				serverWentDown(serv)
			}
			time.Sleep(time.Duration(timeout) * time.Second)
		} else {
			if err == nil {
				serverStartedUp(serv)
			} else {
				serv.IncrementOffTime()
				if shouldReportAgain(serv) {
					notification.ServerStillIsDown(serv)
				}
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func shouldReportAgain(server server.Server) bool {
	// todo эту уникальную формулу явно стоит переделать
	return server.GetOffTime()&0b111 == 0b111 && (server.GetOffTime() < 30 || rand.Intn(4) == 1)
}

func serverWentDown(watcher server.Server) {
	watcher.SetWorking(false)
	log.Info("Server " + watcher.GetFormattedName() + " not working")
	notification.ServerWentDown(watcher)
}

func serverStartedUp(watcher server.Server) {
	watcher.SetWorking(true)
	log.Info("Server " + watcher.GetFormattedName() + " is working again")
	notification.ServerIsUp(watcher)
}
