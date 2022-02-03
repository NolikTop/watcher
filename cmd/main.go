package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"watcher/pkg/config"
	"watcher/pkg/notification"
	"watcher/pkg/server"
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

	servers, err := getServers(c.Servers)
	if err != nil {
		log.Fatal(err)
	}

	err = checkMethodNamesInServers(servers)
	if err != nil {
		log.Fatal(err)
	}

	runWatchers(servers)

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

func getServers(serversConfigs []*config.ServerConfig) ([]server.Server, error) {
	servers := make([]server.Server, len(serversConfigs))

	for i, serverConfig := range serversConfigs {
		serv, err := server.NewServer(serverConfig)
		if err != nil {
			return nil, err
		}

		servers[i] = serv
	}

	return servers, nil
}

func checkMethodNamesInServers(servers []server.Server) error {
	chats := notification.GetMethods()

	for _, serv := range servers {
		for _, chatName := range serv.GetChats() {
			if _, ok := chats[chatName]; !ok {
				return errUnknownChatName(serv.GetName(), chatName)
			}
		}
	}

	return nil
}

func runWatchers(watchers []server.Server) {
	for _, watcher := range watchers {
		go watch.Watch(watcher)
	}
}
