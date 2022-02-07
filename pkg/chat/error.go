package chat

import "errors"

func errNoFieldInData(fieldName string) error {
	return errors.New("no \"" + fieldName + "\" in notification method's \"data\" field")
}
