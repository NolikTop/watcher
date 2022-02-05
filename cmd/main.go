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

	log.Info("Loading config...")

	configPath := flag.String("config", "no", "path to JSON config")
	flag.Parse()

	c, err := config.ParseConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	notification.Init()

	log.Info("Loading notification methods...")

	err = addNotificationMethods(c.NotificationMethods)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Loading servers...")

	servers, err := getServers(c.Servers)
	if err != nil {
		log.Fatal(err)
	}

	err = checkMethodNamesInServers(servers)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Starting watchers...")

	runWatchers(servers)

	log.Info("Watcher started successfully")

	select {} // возможно тут стоило бы просто ловить сигнал от ос об остановке процесса.. но зачем?
}

func addNotificationMethods(methodsConfigs []*config.NotificationMethodConfig) error {
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

func runWatchers(servers []server.Server) {
	for _, serv := range servers {
		logrus.Info("Watching for " + serv.GetFormattedName())
		go watch.Watch(serv)
	}
}
