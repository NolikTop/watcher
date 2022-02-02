package serverwatcher

import (
	rcon "github.com/micvbang/pocketmine-rcon"
)

type RconServerWatcher struct {
	*WatcherBase
	conn     *rcon.Connection
	command  string
	password string
}

func (watcher *RconServerWatcher) Init(data map[string]interface{}) error {

	if command, ok := data["command"]; ok {
		watcher.command = command.(string)
	} else {
		return errNoFieldInData("command")
	}

	if password, ok := data["password"]; ok {
		watcher.password = password.(string)
	} else {
		return errNoFieldInData("password")
	}

	return nil
}

func (watcher *RconServerWatcher) CheckConnection() (err error) {
	if watcher.conn == nil {
		watcher.conn, err = rcon.NewConnection(watcher.serverAddr, watcher.password)
	}

	_, err = watcher.conn.SendCommand(watcher.command)

	if err != nil {
		watcher.conn = nil
	}

	return err
}
