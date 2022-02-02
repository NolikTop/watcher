package main

import "errors"

func errUnknownChatName(serverName, chatName string) error {
	return errors.New("Server " + serverName + " has unknown chat \"" + chatName + "\" in \"chats\" field")
}
