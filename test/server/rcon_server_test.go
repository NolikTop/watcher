package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"watcher/pkg/config"
	"watcher/pkg/serverwatcher"
)

// к rcon нормальные тесты сделать не выйдет потому что рабочих rcon серверов на go я просто не нашел
// есть 3 варианта протестить, и все мне не нравятся.
// 1 - протестить руками (самому запустить сервак локально и к нему попытаться подрубиться) (что в целом я и сделал)
// 2 - запускать каждый раз для теста какой-нибудь pmmp/nukkit и чекать подключение к ним (слишком жирно для тестов)
// 3 - самому написать (вот только зачем это делать ради тестов, да и сам rcon сервер по факту никому не нужен)

func TestRconServerWatcherNoData(t *testing.T) {
	cfg := &config.ServerConfig{
		Name:         "Test RCON ServerWatcher",
		Addr:         "127.0.0.1" + rconServerAddress,
		Protocol:     "rcon",
		MentionsText: "@all",
		Data:         nil,
	}

	_, err := serverwatcher.NewServer(cfg)
	assert.Error(t, err)
}

func TestRconServerWatcherWithEmptyData(t *testing.T) {
	data := make(map[string]interface{})

	cfg := &config.ServerConfig{
		Name:         "Test RCON ServerWatcher",
		Addr:         "127.0.0.1" + rconServerAddress,
		Protocol:     "rcon",
		MentionsText: "@all",
		Data:         data,
	}

	_, err := serverwatcher.NewServer(cfg)
	assert.Error(t, err)
}

func TestRconServerWatcherWithCommandAndNoPassword(t *testing.T) {
	data := make(map[string]interface{})
	data["command"] = "test"

	cfg := &config.ServerConfig{
		Name:         "Test RCON ServerWatcher",
		Addr:         "127.0.0.1" + rconServerAddress,
		Protocol:     "rcon",
		MentionsText: "@all",
		Data:         data,
	}

	_, err := serverwatcher.NewServer(cfg)
	assert.Error(t, err)
}

func TestRconServerWatcherWithPasswordAndNoCommand(t *testing.T) {
	data := make(map[string]interface{})
	data["password"] = "test"

	cfg := &config.ServerConfig{
		Name:         "Test RCON ServerWatcher",
		Addr:         "127.0.0.1" + rconServerAddress,
		Protocol:     "rcon",
		MentionsText: "@all",
		Data:         data,
	}

	_, err := serverwatcher.NewServer(cfg)
	assert.Error(t, err)
}
