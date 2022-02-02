package notification

import (
	log "github.com/sirupsen/logrus"
	"watcher/pkg/serverwatcher"
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

func ServerWentDown(watcher serverwatcher.ServerWatcher) {
	for _, chat := range watcher.GetChats() {
		method := methods[chat]

		err := method.NotifyServerWentDown(watcher)
		if err != nil {
			log.Error(err)
		}
	}
}

func ServerStillIsDown(watcher serverwatcher.ServerWatcher) {
	for _, chat := range watcher.GetChats() {
		method := methods[chat]

		err := method.NotifyServerStillIsDown(watcher)
		if err != nil {
			log.Error(err)
		}
	}
}

func ServerIsUp(watcher serverwatcher.ServerWatcher) {
	for _, chat := range watcher.GetChats() {
		method := methods[chat]

		err := method.NotifyServerIsUp(watcher)
		if err != nil {
			log.Error(err)
		}
	}
}
