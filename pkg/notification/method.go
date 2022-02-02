package notification

import "watcher/pkg/serverwatcher"

type Method interface {
	GetName() string

	Init(name string, data map[string]interface{}) error

	NotifyServerWentDown(server serverwatcher.ServerWatcher) error
	NotifyServerStillIsDown(server serverwatcher.ServerWatcher) error
	NotifyServerIsUp(server serverwatcher.ServerWatcher) error
}
