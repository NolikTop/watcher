package server

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"watcher/pkg/config"
	"watcher/pkg/server"
)

func TestTcpServerWatcher(t *testing.T) {
	serv, err := makeTcpServerWatcher(tcpServerAddress)
	if err != nil {
		t.Fatal(err)
	}

	assert.Error(t, serv.CheckConnection())

	status := make(chan int)

	go runTcpServer(t, tcpServerAddress, status)

	assert.Equal(t, statusServerIsStarted, <-status)
	if !assert.NoError(t, serv.CheckConnection()) {
		t.Fatal("After the server starts successfully, it should connect without errors")
	}

	status <- statusStopServer

	assert.Equal(t, statusServerIsStopped, <-status)
	assert.Error(t, serv.CheckConnection())
}

func makeTcpServerWatcher(address string) (server.Server, error) {
	cfg := &config.ServerConfig{
		Name:         "Test TCP Server",
		Addr:         "127.0.0.1" + address,
		Protocol:     "tcp",
		MentionsText: "@all",
		Data:         nil,
	}

	return server.NewServer(cfg)
}

func runTcpServer(t *testing.T, address string, status chan int) {
	serv, err := net.Listen("tcp", address)
	if err != nil {
		t.Fatal(err)
	}

	status <- statusServerIsStarted

	_, err = serv.Accept()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, statusStopServer, <-status)

	err = serv.Close()
	if err != nil {
		t.Error(err)
	}

	status <- statusServerIsStopped
}
