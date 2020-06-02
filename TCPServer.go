package main

import "net"

type TCPServer struct {
	*ServerBase
}

func (server *TCPServer) checkConnection() (err error) {
	_, err = net.Dial("tcp", server.Addr)
	if err != nil {
		return
	}

	// мб отправлять какие-нибудь байты?
	return nil
}
