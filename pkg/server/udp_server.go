package server

import (
	"encoding/base64"
	"net"
	"time"
)

type UdpServer struct {
	*Base
	sendBytes []byte
	conn      net.Conn
	buffer    []byte
}

func (s *UdpServer) Init(data map[string]interface{}) error {
	if sendBytesBase64, ok := data["send_bytes_base64"]; ok {
		sendBytes, err := base64.StdEncoding.DecodeString(sendBytesBase64.(string))
		if err != nil {
			return err
		}

		s.sendBytes = sendBytes
	} else {
		return errNoFieldInData("send_bytes_base64")
	}

	return nil
}

func (s *UdpServer) CheckConnection() (err error) {
	if s.conn == nil {
		s.conn, err = net.Dial("udp", s.serverAddr)
		s.buffer = make([]byte, 64)
	}

	_, err = s.conn.Write(s.sendBytes)
	if err != nil {
		return
	}

	deadline := time.Now().Add(2 * time.Second)
	err = s.conn.SetReadDeadline(deadline)
	if err != nil {
		return
	}

	for i := range s.buffer {
		s.buffer[i] = 0
	}

	n, err := s.conn.Read(s.buffer)
	if err != nil {
		return
	}

	if n == 0 {
		return errNoBytesReceived
	}

	return nil
}
