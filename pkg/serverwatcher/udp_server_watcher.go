package serverwatcher

import (
	"encoding/base64"
	"net"
	"time"
)

type UdpServerWatcher struct {
	*WatcherBase
	sendBytes []byte
	conn      net.Conn
	buffer    []byte
}

func (w *UdpServerWatcher) Init(data map[string]interface{}) error {
	if sendBytesBase64, ok := data["send_bytes_base64"]; ok {
		sendBytes, err := base64.StdEncoding.DecodeString(sendBytesBase64.(string))
		if err != nil {
			return err
		}

		w.sendBytes = sendBytes
	} else {
		return errNoFieldInData("send_bytes_base64")
	}

	return nil
}

func (w *UdpServerWatcher) CheckConnection() (err error) {
	if w.conn == nil {
		w.conn, err = net.Dial("udp", w.serverAddr)
		w.buffer = make([]byte, 64)
	}

	_, err = w.conn.Write(w.sendBytes)
	if err != nil {
		return
	}

	deadline := time.Now().Add(2 * time.Second)
	err = w.conn.SetReadDeadline(deadline)
	if err != nil {
		return
	}

	for i := range w.buffer {
		w.buffer[i] = 0
	}

	n, err := w.conn.Read(w.buffer)
	if err != nil {
		return
	}

	if n == 0 {
		return errNoBytesReceived
	}

	return nil
}
