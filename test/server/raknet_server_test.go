package server

import (
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/stretchr/testify/assert"
	"testing"
	"watcher/pkg/config"
	"watcher/pkg/serverwatcher"
)

func TestRaknetServerWatcher(t *testing.T) {
	serv, err := makeRaknetServerWatcher(raknetServerAddress)
	if err != nil {
		t.Fatal(err)
	}

	assert.Error(t, serv.CheckConnection())

	status := make(chan int)

	go runRaknetServer(t, raknetServerAddress, status)

	assert.Equal(t, statusServerIsStarted, <-status)
	if !assert.NoError(t, serv.CheckConnection()) {
		t.Fatal("After the server starts successfully, it should connect without errors")
	}

	status <- statusStopServer

	assert.Equal(t, statusServerIsStopped, <-status)
	assert.Error(t, serv.CheckConnection())
}

func makeRaknetServerWatcher(address string) (serverwatcher.ServerWatcher, error) {
	cfg := &config.ServerConfig{
		Name:         "Test Raknet ServerWatcher",
		Addr:         "127.0.0.1" + address,
		Protocol:     "raknet",
		MentionsText: "@all",
		Data:         nil,
	}

	return serverwatcher.NewServer(cfg)
}

func runRaknetServer(t *testing.T, address string, status chan int) {
	serv, err := minecraft.ListenConfig{}.Listen("raknet", address)
	if err != nil {
		t.Fatal(err)
	}

	status <- statusServerIsStarted

	assert.Equal(t, statusStopServer, <-status)

	err = serv.Close()
	if err != nil {
		t.Error(err)
	}

	status <- statusServerIsStopped
}
