package serverwatcher

import (
	"net"
)

type TcpServerWatcher struct {
	*WatcherBase
}

func (w *TcpServerWatcher) Init(data map[string]interface{}) error {
	return nil
}

func (w *TcpServerWatcher) CheckConnection() (err error) {
	conn, err := net.Dial("tcp", w.serverAddr)
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
