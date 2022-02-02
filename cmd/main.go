package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"watcher/pkg/config"
)

func main() {
	logger := logrus.New()

	configPath := flag.String("config", "no", "path to JSON config")
	flag.Parse()

	c, err := config.ParseConfig(*configPath)
	if err != nil {
		logger.Fatal(err)
	}

	select {}
}
