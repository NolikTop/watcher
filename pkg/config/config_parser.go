package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

func ParseConfig(configPath string) (*Config, error) {
	file, err := verifyConfigExistsAndGetContents(configPath)
	if err != nil {
		return nil, err
	}

	config := &RawConfig{}
	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	err = verifyAllRequiredFields(config)
	if err != nil {
		return nil, err
	}

	chats, err := parseAllChats(config)
	if err != nil {
		return nil, err
	}

	servers, err := parseAllServers(config)
	if err != nil {
		return nil, err
	}

	return &Config{
		Chats:   chats,
		Servers: servers,
	}, nil
}

var errNoPathSpecified = errors.New("Specify JSON config path using -config=path ")

func verifyConfigExistsAndGetContents(configPath string) ([]byte, error) {
	if configPath == "no" {
		return nil, errNoPathSpecified
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func verifyAllRequiredFields(config *RawConfig) error {
	if len(config.Chats) == 0 {
		return errFieldIsNotProvided("notification_methods")
	}

	if len(config.Servers) == 0 {
		return errFieldIsNotProvided("servers")
	}

	return nil
}

func parseAllChats(config *RawConfig) ([]*ChatConfig, error) {
	chats := make([]*ChatConfig, len(config.Chats))

	for id, chatData := range config.Chats {
		chatConfig, err := parseChat(id, chatData)
		if err != nil {
			return nil, err
		}

		chats[id] = chatConfig
	}

	return chats, nil
}

func parseChat(id int, chatData *RawChatConfig) (*ChatConfig, error) {
	err := verifyChatRequiredFields(id, chatData)
	if err != nil {
		return nil, err
	}

	return &ChatConfig{
		Name:   *chatData.Name,
		Method: *chatData.Method,
		Data:   *chatData.Data,
	}, nil
}

func verifyChatRequiredFields(id int, chatData *RawChatConfig) error {
	if chatData.Name == nil {
		serverName := "#" + strconv.Itoa(id) // если имя метода не указано, то в ошибке укажем его порядковый номер в конфиге
		return errChatHasNotField(serverName, "name")
	}

	serverName := *chatData.Name

	if chatData.Method == nil {
		return errChatHasNotField(serverName, "method")
	}

	if chatData.Data == nil {
		data := make(map[string]interface{})
		chatData.Data = &data
	}

	return nil
}

func parseAllServers(config *RawConfig) ([]*ServerConfig, error) {
	servers := make([]*ServerConfig, len(config.Servers))

	for id, serverData := range config.Servers {
		serverConfig, err := parseConfigServer(id, serverData)
		if err != nil {
			return nil, err
		}

		servers[id] = serverConfig
	}

	return servers, nil
}

func parseConfigServer(id int, serverData *RawServerConfig) (*ServerConfig, error) {
	err := verifyServerRequiredFields(id, serverData)
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		Name:         *serverData.Name,
		Addr:         *serverData.Addr,
		Protocol:     *serverData.Protocol,
		Chats:        *serverData.Chats,
		MentionsText: *serverData.MentionsText,
		Data:         *serverData.Data,
	}, nil
}

func verifyServerRequiredFields(id int, serverData *RawServerConfig) error {
	if serverData.Name == nil {
		serverName := "#" + strconv.Itoa(id) // если имя сервера не указано, то в ошибке укажем его порядковый номер в конфиге
		return errServerHasNotField(serverName, "name")
	}

	serverName := *serverData.Name

	if serverData.Addr == nil {
		return errServerHasNotField(serverName, "addr")
	}

	if serverData.Protocol == nil {
		return errServerHasNotField(serverName, "server")
	}

	if serverData.Chats == nil {
		return errServerHasNotField(serverName, "chats")
	}

	if serverData.MentionsText == nil {
		return errServerHasNotField(serverName, "mentions_text")
	}

	if serverData.Data == nil {
		data := make(map[string]interface{})
		serverData.Data = &data
	}

	return nil
}
