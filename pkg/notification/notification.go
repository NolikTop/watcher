package notification

import (
	log "github.com/sirupsen/logrus"
	"watcher/pkg/server"
)

var methods map[string]Method

func Init() {
	methods = make(map[string]Method)
}

func GetMethods() map[string]Method {
	return methods
}

func Add(notificationMethod Method) error {
	name := notificationMethod.GetName()
	if _, ok := methods[name]; ok {
		return errChatWithThisNameAlreadyExists(name)
	}

	methods[name] = notificationMethod

	return nil
}

func ServerWentDown(serv server.Server) {
	for _, chat := range serv.GetChats() {
		method := methods[chat]

		err := method.NotifyServerWentDown(serv)
		if err != nil {
			log.Error(err)
		}
	}
}

func ServerStillIsDown(serv server.Server) {
	for _, chat := range serv.GetChats() {
		method := methods[chat]

		err := method.NotifyServerStillIsDown(serv)
		if err != nil {
			log.Error(err)
		}
	}
}

func ServerIsUp(serv server.Server) {
	for _, chat := range serv.GetChats() {
		method := methods[chat]

		err := method.NotifyServerIsUp(serv)
		if err != nil {
			log.Error(err)
		}
	}
}
