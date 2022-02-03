package server

import (
	"net"
)

type TcpServer struct {
	*Base
}

func (s *TcpServer) Init(data map[string]interface{}) error {
	return nil
}

func (s *TcpServer) CheckConnection() (err error) {
	conn, err := net.Dial("tcp", s.serverAddr)
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
