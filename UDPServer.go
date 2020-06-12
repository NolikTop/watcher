package main

import (
	"encoding/base64"
	"errors"
	"net"
	"time"
)

type UDPServer struct {
	*ServerBase
	startBytes []byte
	conn       net.Conn
	tempByteAr []byte
}

func (server *UDPServer) checkConnection() (err error) {
	if server.conn == nil {
		server.conn, err = net.Dial("udp", server.Addr)
		server.tempByteAr = make([]byte, 64)
	}

	_, err = server.conn.Write(server.startBytes)
	if err != nil {
		return
	}

	deadline := time.Now().Add(2 * time.Second)
	err = server.conn.SetReadDeadline(deadline)
	if err != nil {
		return
	}

	for i := range server.tempByteAr {
		server.tempByteAr[i] = 0
	}
	n, err := server.conn.Read(server.tempByteAr)
	if err != nil {
		return
	}

	if n == 0 {
		return errors.New("no bytes received: " + base64.StdEncoding.EncodeToString(server.tempByteAr))
	}

	return nil
}
