package notification

import "watcher/pkg/server"

type Method interface {
	GetName() string

	Init(name string, data map[string]interface{}) error

	NotifyServerWentDown(server server.Server) error
	NotifyServerStillIsDown(server server.Server) error
	NotifyServerIsUp(server server.Server) error
}
