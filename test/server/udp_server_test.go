package server

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"watcher/pkg/config"
	"watcher/pkg/serverwatcher"
)

func TestUdpServerWatcherNoData(t *testing.T) {
	cfg := &config.ServerConfig{
		Name:         "Test TCP ServerWatcher",
		Addr:         "127.0.0.1" + udpServerAddress,
		Protocol:     "udp",
		MentionsText: "@all",
		Data:         nil,
	}

	_, err := serverwatcher.NewServer(cfg)
	assert.Error(t, err)
}

func TestUdpServerWatcherWithDataAndNoSendBytes(t *testing.T) {
	data := make(map[string]interface{})

	cfg := &config.ServerConfig{
		Name:         "Test TCP ServerWatcher",
		Addr:         "127.0.0.1" + udpServerAddress,
		Protocol:     "udp",
		MentionsText: "@all",
		Data:         data,
	}

	_, err := serverwatcher.NewServer(cfg)
	assert.Error(t, err)
}

func TestUdpServerWatcherWithDataAndWrongBase64SendBytes(t *testing.T) {
	wrongBase64 := "123456789!@#$%^&*"

	data := make(map[string]interface{})
	data["send_bytes_base64"] = wrongBase64

	cfg := &config.ServerConfig{
		Name:         "Test TCP ServerWatcher",
		Addr:         "127.0.0.1" + udpServerAddress,
		Protocol:     "udp",
		MentionsText: "@all",
		Data:         data,
	}

	_, err := serverwatcher.NewServer(cfg)
	assert.Error(t, err)
}

func TestUdpServerWatcher(t *testing.T) {
	serv, err := makeUdpServerWatcher(udpServerAddress)
	if err != nil {
		t.Fatal(err)
	}

	assert.Error(t, serv.CheckConnection())

	status := make(chan int)

	go runUdpServer(t, udpServerAddress, status)

	assert.Equal(t, statusServerIsStarted, <-status)
	if !assert.NoError(t, serv.CheckConnection()) {
		t.Fatal("After the server starts successfully, it should connect without errors")
	}

	status <- statusStopServer

	assert.Equal(t, statusServerIsStopped, <-status)
	assert.Error(t, serv.CheckConnection())
}

func makeUdpServerWatcher(address string) (serverwatcher.ServerWatcher, error) {
	sendBytes := "1234567890"

	data := make(map[string]interface{})
	data["send_bytes_base64"] = base64.StdEncoding.EncodeToString([]byte(sendBytes))

	cfg := &config.ServerConfig{
		Name:         "Test UDP ServerWatcher",
		Addr:         "127.0.0.1" + address,
		Protocol:     "udp",
		MentionsText: "@all",
		Data:         data,
	}

	return serverwatcher.NewServer(cfg)
}

func runUdpServer(t *testing.T, address string, status chan int) {
	serv, err := net.ListenPacket("udp", address)
	if err != nil {
		t.Fatal(err)
	}

	status <- statusServerIsStarted

	buf := make([]byte, 1024)

	n, addr, err := serv.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	n, err = serv.WriteTo(buf[:n], addr) // обратно отсылаем те же байты, почему бы и нет ¯\_(ツ)_/¯
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
