package server

import "errors"

var errNoBytesReceived = errors.New("no bytes received")

func errNoFieldInData(fieldName string) error {
	return errors.New("No \"" + fieldName + "\" in server's \"data\" field")
}
