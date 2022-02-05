package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
	"watcher/pkg/config"
	"watcher/pkg/server"
)

func TestHttpServerWatcher(t *testing.T) {
	serv, err := makeHttpServerWatcher(httpServerAddress)
	if err != nil {
		t.Fatal(err)
	}

	assert.Error(t, serv.CheckConnection())

	status := make(chan int)

	go runHttpServer(t, httpServerAddress, status)

	assert.Equal(t, statusServerIsStarted, <-status)
	time.Sleep(time.Millisecond * 50) // http сервер мгновенно не запускается, более того в стандартной реализации нет возможности узнать когда именно он запустился

	if !assert.NoError(t, serv.CheckConnection()) {
		t.Fatal("After the server starts successfully, it should connect without errors")
	}

	status <- statusStopServer

	assert.Equal(t, statusServerIsStopped, <-status)
	assert.Error(t, serv.CheckConnection())
}

func makeHttpServerWatcher(address string) (server.Server, error) {
	cfg := &config.ServerConfig{
		Name:         "Test HTTP Server",
		Addr:         "127.0.0.1" + address,
		Protocol:     "http",
		MentionsText: "@all",
		Data:         nil,
	}

	return server.NewServer(cfg)
}

func runHttpServer(t *testing.T, address string, status chan int) {
	serv := &http.Server{
		Addr: address,
	}

	http.HandleFunc("/", handler)

	status <- statusServerIsStarted

	go func() {
		assert.Equal(t, statusStopServer, <-status)

		err := serv.Close()
		if err != nil {
			t.Error(err)
		}

		status <- statusServerIsStopped
	}()

	err := serv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		t.Error(err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello")
	if err != nil {
		log.Error(err)
	}
}
