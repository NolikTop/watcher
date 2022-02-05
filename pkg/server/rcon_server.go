package server

import (
	rcon "github.com/micvbang/pocketmine-rcon"
)

type RconServer struct {
	*Base
	conn     *rcon.Connection
	command  string
	password string
}

func (s *RconServer) Init(data map[string]interface{}) error {

	if command, ok := data["command"]; ok {
		s.command = command.(string)
	} else {
		return errNoFieldInData("command")
	}

	if password, ok := data["password"]; ok {
		s.password = password.(string)
	} else {
		return errNoFieldInData("password")
	}

	return nil
}

func (s *RconServer) CheckConnection() (err error) {
	if s.conn == nil {
		s.conn, err = rcon.NewConnection(s.serverAddr, s.password)
		if err != nil {
			return err
		}
	}

	_, err = s.conn.SendCommand(s.command)

	if err != nil {
		s.conn = nil
	}

	return err
}
