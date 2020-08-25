package main

import (
	"encoding/base64"
	"errors"
	"github.com/GreenWix/binary"
	"net"
	"time"
)

type MinecraftServer struct {
	*ServerBase
	conn       net.Conn
	tempByteAr []byte
}

func (server *MinecraftServer) checkConnection() (err error) {
	if server.conn == nil {
		server.conn, err = net.Dial("udp", server.Addr)
		server.tempByteAr = make([]byte, 64)
	}

	w := binary.AcquireWriter(1 + 8)
	w.WriteByte(0x01)                                                                       // ID_UNCONNECTED_PING
	w.WriteSignedLong(time.Now().Unix() * 1000)                                             // sendPingTime
	w.Write(16, []byte("\x00\xff\xff\x00\xfe\xfe\xfe\xfe\xfd\xfd\xfd\xfd\x12\x34\x56\x78")) // magic
	w.WriteSignedLong(0)                                                                    // client Id

	startBytes := w.Buffer()
	binary.ReleaseWriter(w)

	_, err = server.conn.Write(startBytes)
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
