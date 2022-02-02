package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"watcher/pkg/config"
	"watcher/pkg/notification"
	"watcher/pkg/serverwatcher"
	"watcher/pkg/watch"
)

func main() {
	log := logrus.New()

	configPath := flag.String("config", "no", "path to JSON config")
	flag.Parse()

	c, err := config.ParseConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = addMethods(c.NotificationMethods)
	if err != nil {
		log.Fatal(err)
	}

	watchers, err := getWatchers(c.Servers)
	if err != nil {
		log.Fatal(err)
	}

	err = checkMethodNamesInServers(watchers)
	if err != nil {
		log.Fatal(err)
	}

	runWatchers(watchers)

	select {}
}

func addMethods(methodsConfigs []*config.NotificationMethodConfig) error {
	for _, methodConfig := range methodsConfigs {
		method, err := notification.NewMethod(methodConfig)
		if err != nil {
			return err
		}

		err = notification.Add(method)
		if err != nil {
			return err
		}
	}

	return nil
}

func getWatchers(serversConfigs []*config.ServerConfig) ([]serverwatcher.ServerWatcher, error) {
	watchers := make([]serverwatcher.ServerWatcher, len(serversConfigs))

	for _, serverConfig := range serversConfigs {
		watcher, err := serverwatcher.NewServer(serverConfig)
		if err != nil {
			return nil, err
		}

		watchers = append(watchers, watcher)
	}

	return watchers, nil
}

func checkMethodNamesInServers(watchers []serverwatcher.ServerWatcher) error {
	chats := notification.GetMethods()

	for _, watcher := range watchers {
		for _, chatName := range watcher.GetChats() {
			if _, ok := chats[chatName]; !ok {
				return errUnknownChatName(watcher.GetName(), chatName)
			}
		}
	}

	return nil
}

func runWatchers(watchers []serverwatcher.ServerWatcher) {
	for _, watcher := range watchers {
		go watch.Watch(watcher)
	}
}
