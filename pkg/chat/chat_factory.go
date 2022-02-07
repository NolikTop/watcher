package chat

import (
	"errors"
	"github.com/NolikTop/watcher/pkg/config"
)

func NewChat(config *config.ChatConfig) (Chat, error) {
	var cht Chat

	method := config.Method
	switch method {
	case "vk":
		cht = &Vk{}
	default:
		return nil, errUnknownMethod(method)
	}

	err := cht.Init(config.Name, config.Data)
	if err != nil {
		return nil, err
	}

	return cht, nil
}

func errUnknownMethod(method string) error {
	return errors.New("unknown method \"" + method + "\"")
}
