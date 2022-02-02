package notification

import "errors"

func errNoFieldInData(fieldName string) error {
	return errors.New("No \"" + fieldName + "\" in notification method's \"data\" field")
}

func errChatWithThisNameAlreadyExists(chatName string) error {
	return errors.New("Chat with name \"" + chatName + "\" already exists")
}
