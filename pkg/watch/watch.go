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

		if serv.IsMarkedAsWorking() {
			if err != nil {
				serverWentDown(serv, err)
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

func serverWentDown(server server.Server, err error) {
	server.MarkIsWorking(false)
	log.WithField("server", server.GetFormattedName()).WithField("error", err).Info("Server went down")
	notification.ServerWentDown(server, err)
}

func serverStartedUp(server server.Server) {
	server.MarkIsWorking(true)
	log.WithField("server", server.GetFormattedName()).Info("Server is working again")
	notification.ServerIsUp(server)
}
