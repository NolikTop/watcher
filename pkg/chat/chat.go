package chat

import "github.com/NolikTop/watcher/pkg/server"

type Chat interface {
	GetName() string

	Init(name string, data map[string]interface{}) error

	NotifyServerWentDown(server server.Server, err error) error
	NotifyServerStillIsDown(server server.Server) error
	NotifyServerIsUp(server server.Server) error
}
