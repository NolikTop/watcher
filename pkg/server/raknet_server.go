package server

import (
	"github.com/GreenWix/binary"
	"net"
	"time"
)

type RaknetServer struct {
	*Base
	conn   net.Conn
	buffer []byte
}

func (s *RaknetServer) Init(data map[string]interface{}) error {
	return nil
}

const UnconnectedPingSize = 1 + 8 + 16 + 8 // byte + long + magic bytes (16) + long

func (s *RaknetServer) CheckConnection() error {
	if s.conn == nil {
		err := s.setRaknetConnectionAndBuffer()
		if err != nil {
			return err
		}
	}

	unconnectedPingBytes := makeUnconnectedPingPacket()
	_, err := s.conn.Write(unconnectedPingBytes)
	if err != nil {
		return err
	}

	deadline := time.Now().Add(2 * time.Second)
	err = s.conn.SetReadDeadline(deadline)
	if err != nil {
		return err
	}

	for i := range s.buffer {
		s.buffer[i] = 0
	}

	n, err := s.conn.Read(s.buffer)
	if err != nil {
		return err
	}

	if n == 0 {
		return errNoBytesReceived
	}

	return nil
}

func (s *RaknetServer) setRaknetConnectionAndBuffer() error {
	var err error

	s.conn, err = net.Dial("udp", s.serverAddr)
	if err != nil {
		return err
	}

	s.buffer = make([]byte, 64)

	return nil
}

func makeUnconnectedPingPacket() []byte {
	w := binary.AcquireWriter(UnconnectedPingSize)
	w.WriteByte(0x01)                                                                       // ID_UNCONNECTED_PING
	w.WriteSignedLong(time.Now().Unix() * 1000)                                             // sendPingTime
	w.Write(16, []byte("\x00\xff\xff\x00\xfe\xfe\xfe\xfe\xfd\xfd\xfd\xfd\x12\x34\x56\x78")) // magic (эта часть буквально называется magic)
	w.WriteSignedLong(0)                                                                    // client Id

	defer binary.ReleaseWriter(w)
	return w.Buffer()
}
