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

func (watcher *UdpServerWatcher) Init(data map[string]interface{}) error {
	if sendBytesBase64, ok := data["send_bytes_base64"]; ok {
		sendBytes, err := base64.StdEncoding.DecodeString(sendBytesBase64.(string))
		if err != nil {
			return err
		}

		watcher.sendBytes = sendBytes
	} else {
		return errNoFieldInData("send_bytes_base64")
	}

	return nil
}

func (watcher *UdpServerWatcher) CheckConnection() (err error) {
	if watcher.conn == nil {
		watcher.conn, err = net.Dial("udp", watcher.serverAddr)
		watcher.buffer = make([]byte, 64)
	}

	_, err = watcher.conn.Write(watcher.sendBytes)
	if err != nil {
		return
	}

	deadline := time.Now().Add(2 * time.Second)
	err = watcher.conn.SetReadDeadline(deadline)
	if err != nil {
		return
	}

	for i := range watcher.buffer {
		watcher.buffer[i] = 0
	}

	n, err := watcher.conn.Read(watcher.buffer)
	if err != nil {
		return
	}

	if n == 0 {
		return errNoBytesReceived
	}

	return nil
}
