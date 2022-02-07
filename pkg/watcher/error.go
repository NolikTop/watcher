package watcher

import "errors"

func errMethodWithThisNameAlreadyExists(chatName string) error {
	return errors.New("method with name \"" + chatName + "\" already exists")
}

func errUnknownChatName(serverName, chatName string) error {
	return errors.New("Server " + serverName + " has unknown chat \"" + chatName + "\" in \"chats\" field")
}

var errWatcherCantBeChanged = errors.New("watcher cant be changed after start")
