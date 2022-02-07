package watcher

import "errors"

func errChatWithThisNameAlreadyExists(chatName string) error {
	return errors.New("chat with name \"" + chatName + "\" already exists")
}

func errUnknownChatName(serverName, chatName string) error {
	return errors.New("server " + serverName + " has unknown chat \"" + chatName + "\" in \"chats\" field")
}

var errWatcherCantBeChanged = errors.New("watcher cant be changed after start")
