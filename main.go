package main

import (
	"flag"
	"github.com/NolikTop/watcher/pkg/config"
	"github.com/NolikTop/watcher/pkg/watcher"
	"github.com/sirupsen/logrus"
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

	w := watcher.New()

	log.Info("Adding data from config...")
	err = w.Load(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Starting watchers...")
	err = w.Start()
	if err != nil {
		log.Fatal(err)
	}

	select {} // возможно тут стоило бы просто ловить сигнал от ос об остановке процесса.. но зачем?
}
