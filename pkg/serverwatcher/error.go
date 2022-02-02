package serverwatcher

import "errors"

var errNoBytesReceived = errors.New("no bytes received")

func errNoFieldInData(fieldName string) error {
	return errors.New("No \"" + fieldName + "\" in serverwatcher's \"data\" field")
}
