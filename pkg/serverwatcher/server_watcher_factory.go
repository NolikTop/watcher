package serverwatcher

import (
	"errors"
	"watcher/pkg/config"
)

func newBase(config *config.ServerConfig) *WatcherBase {
	return &WatcherBase{
		name:         config.Name,
		serverAddr:   config.Addr,
		working:      true,
		offTime:      0,
		mentionsText: config.MentionsText,
	}
}

func NewServer(config *config.ServerConfig) (ServerWatcher, error) {
	var server ServerWatcher

	base := newBase(config)

	protocol := config.Protocol
	switch protocol {
	case "udp":
		server = &UdpServerWatcher{WatcherBase: base}
	case "tcp":
		server = &TcpServerWatcher{WatcherBase: base}
	case "raknet":
		server = &RakNetServerWatcher{WatcherBase: base}
	case "minecraft":
		return nil, errProtocolNotImplemented(protocol)
	case "rcon":
		server = &RconServerWatcher{WatcherBase: base}
	default:
		return nil, errUnknownProtocol(protocol)
	}

	err := server.Init(config.Data)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func errProtocolNotImplemented(serverProtocol string) error {
	return errors.New("Protocol \"" + serverProtocol + "\" is not implemented yet")
}

func errUnknownProtocol(serverProtocol string) error {
	return errors.New("Unknown serverwatcher \"" + serverProtocol + "\"")
}
