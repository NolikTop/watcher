package chat

import "errors"

func errNoFieldInData(fieldName string) error {
	return errors.New("no \"" + fieldName + "\" in chat's \"data\" field")
}
