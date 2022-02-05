package server

import (
	"errors"
	"watcher/pkg/config"
)

func newBase(config *config.ServerConfig) *Base {
	return &Base{
		name:         config.Name,
		serverAddr:   config.Addr,
		working:      true,
		offTime:      0,
		chats:        config.Chats,
		mentionsText: config.MentionsText,
		protocol:     config.Protocol,
	}
}

func NewServer(config *config.ServerConfig) (Server, error) {
	var server Server

	base := newBase(config)

	protocol := config.Protocol
	switch protocol {
	case "udp":
		server = &UdpServer{Base: base}
	case "tcp":
		server = &TcpServer{Base: base}
	case "raknet":
		server = &RaknetServer{Base: base}
	case "minecraft":
		return nil, errProtocolNotImplemented(protocol)
	case "rcon":
		server = &RconServer{Base: base}
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
	return errors.New("Unknown server \"" + serverProtocol + "\"")
}
