package main

import "net"

type TCPServer struct {
	*ServerBase
}

func (server *TCPServer) checkConnection() (err error) {
	conn, err := net.Dial("tcp", server.Addr)
	if err != nil {
		return
	}

	err = conn.Close()
	if err != nil {
		return
	}

	// мб отправлять какие-нибудь байты?
	return nil
}
