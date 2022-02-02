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

	servers, err := parseAllServers(config)

	return &Config{
		VKToken:  *config.VKToken,
		Time:     *config.Time,
		VKChatId: *config.VKChatId,
		Servers:  servers,
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
	if config.VKToken == nil {
		return errFieldIsNotProvided("vk_token")
	}

	if config.VKChatId == nil {
		return errFieldIsNotProvided("vk_chat_id")
	}

	if config.Time == nil {
		return errFieldIsNotProvided("time")
	}

	if len(config.Servers) == 0 {
		return errFieldIsNotProvided("servers")
	}

	return nil
}

func parseAllServers(config *RawConfig) ([]*ServerConfig, error) {
	servers := make([]*ServerConfig, len(config.Servers))
	for id, server := range config.Servers {
		serverConfig, err := parseConfigServer(id, server)
		if err != nil {
			return nil, err
		}

		servers[id] = serverConfig
	}

	return servers, nil
}

func parseConfigServer(id int, server *RawServerConfig) (*ServerConfig, error) {
	err := verifyServerRequiredFields(id, server)
	if err != nil {
		return nil, err
	}

	var serverData map[string]interface{}
	if server.Data == nil {
		serverData = make(map[string]interface{})
	} else {
		serverData = *server.Data
	}

	return &ServerConfig{
		Name:         *server.Name,
		Addr:         *server.Addr,
		Protocol:     *server.Protocol,
		MentionsText: *server.MentionsText,
		Data:         serverData,
	}, nil
}

func verifyServerRequiredFields(id int, server *RawServerConfig) error {
	if server.Name == nil {
		serverName := "#" + strconv.Itoa(id) // если имя сервера не указано, то в ошибке укажем его порядковый номер в конфиге
		return errServerHasNotField(serverName, "name")
	}

	serverName := *server.Name

	if server.Addr == nil {
		return errServerHasNotField(serverName, "addr")
	}

	if server.Protocol == nil {
		return errServerHasNotField(serverName, "serverwatcher")
	}

	if server.MentionsText == nil {
		return errServerHasNotField(serverName, "mentions_text")
	}

	return nil
}

func errFieldIsNotProvided(fieldName string) error {
	return errors.New("No \"" + fieldName + "\" field provided in config ")
}

func errServerHasNotField(serverName, fieldName string) error {
	return errors.New("Server " + serverName + " hasn't \"" + fieldName + "\" field")
}
