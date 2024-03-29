package config

import "errors"

func errFieldIsNotProvided(fieldName string) error {
	return errors.New("No \"" + fieldName + "\" field provided in config ")
}

func errChatHasNotField(chatName, fieldName string) error {
	return errors.New("Chat " + chatName + " hasn't \"" + fieldName + "\" field")
}

func errServerHasNotField(serverName, fieldName string) error {
	return errors.New("Server " + serverName + " hasn't \"" + fieldName + "\" field")
}
