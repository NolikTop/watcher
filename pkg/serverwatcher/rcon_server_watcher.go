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

func (w *RconServerWatcher) Init(data map[string]interface{}) error {

	if command, ok := data["command"]; ok {
		w.command = command.(string)
	} else {
		return errNoFieldInData("command")
	}

	if password, ok := data["password"]; ok {
		w.password = password.(string)
	} else {
		return errNoFieldInData("password")
	}

	return nil
}

func (w *RconServerWatcher) CheckConnection() (err error) {
	if w.conn == nil {
		w.conn, err = rcon.NewConnection(w.serverAddr, w.password)
	}

	_, err = w.conn.SendCommand(w.command)

	if err != nil {
		w.conn = nil
	}

	return err
}
