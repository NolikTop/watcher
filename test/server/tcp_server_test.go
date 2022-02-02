package server

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"watcher/pkg/config"
	"watcher/pkg/serverwatcher"
)

func TestTcpServer(t *testing.T) {
	address := ":1525"

	serv, err := makeTcpServerWatcher(address)
	if err != nil {
		t.Fatal(err)
	}

	assert.Error(t, serv.CheckConnection())

	messageFromServer := make(chan int, 1)
	messageToServer := make(chan int, 1)

	go runTcpServer(t, address, messageFromServer, messageToServer)

	assert.Equal(t, statusServerIsStarted, <-messageFromServer)

	if !assert.NoError(t, serv.CheckConnection()) {
		t.Fatal("After the server starts successfully, it should connect without errors")
	}

	messageToServer <- statusStopServer
	assert.Equal(t, statusServerIsStopped, <-messageFromServer)

	assert.Error(t, serv.CheckConnection())
}

func makeTcpServerWatcher(address string) (serverwatcher.ServerWatcher, error) {
	cfg := &config.ServerConfig{
		Name:         "Test TCP ServerWatcher",
		Addr:         "127.0.0.1" + address,
		Protocol:     "tcp",
		MentionsText: "@all",
		Data:         nil,
	}

	return serverwatcher.NewServer(cfg)
}

func runTcpServer(t *testing.T, address string, messageFromServer chan int, messageToServer chan int) {
	ln, _ := net.Listen("tcp", address)

	messageFromServer <- statusServerIsStarted

	_, err := ln.Accept()
	if err != nil {
		t.Fatal("Couldnt start tcp server", err)
	}

	assert.Equal(t, statusStopServer, <-messageToServer)

	err = ln.Close()
	if err != nil {
		t.Error(err)
	}

	messageFromServer <- statusServerIsStopped
}
