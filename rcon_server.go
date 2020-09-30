package main

import (
	rcon "github.com/micvbang/pocketmine-rcon"
)

type RconServer struct {
	*ServerBase
	conn    *rcon.Connection
	command string
	pass    string
}

func (server *RconServer) checkConnection() (err error) {
	//todo мб стоит постоянно открывать соединение?

	if server.conn == nil {
		server.conn, err = rcon.NewConnection(server.Addr, server.pass)
	}

	_, err = server.conn.SendCommand(server.command)

	if err != nil {
		server.conn = nil
	}

	return err
}
