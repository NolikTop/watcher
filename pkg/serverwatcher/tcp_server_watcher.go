package serverwatcher

import (
	"net"
)

type TcpServerWatcher struct {
	*WatcherBase
}

func (watcher *TcpServerWatcher) Init(data map[string]interface{}) error {
	return nil
}

func (watcher *TcpServerWatcher) CheckConnection() (err error) {
	conn, err := net.Dial("tcp", watcher.serverAddr)
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
