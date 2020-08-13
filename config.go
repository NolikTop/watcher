package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// Существует только для проверки на наличие required полей в конфиге
type RawConfigStruct struct {
	VKToken  *string            `json:"vk_token"`
	VKChatId *int               `json:"vk_chat_id"`
	Time     *int               `json:"time"`
	Servers  []*RawConfigServer `json:"servers"`
}

type ConfigStruct struct {
	VKToken  string
	VKChatId int
	Time     int
	Servers  []*ConfigServer
}

// Существует только для проверки на наличие required полей в конфиге
type RawConfigServer struct {
	Name             *string `json:"name"`
	Addr             *string `json:"addr"`
	Protocol         *string `json:"protocol"`
	MentionsText     *string `json:"mentions_text"`
	StartBytesBase64 *string `json:"start_bytes_base64"`
}

type ConfigServer struct {
	Name             string
	Addr             string
	Protocol         string
	MentionsText     string
	StartBytesBase64 string
}

func parseConfig(configPath string) *ConfigStruct {
	if configPath == "no" {
		panic("Specify JSON config path using -config=path")
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("Couldn't read config file: " + err.Error())
	}

	config := &RawConfigStruct{}

	err = json.Unmarshal(file, config)
	if err != nil {
		panic("Couldn't read config file: " + err.Error())
	}

	if config.VKToken == nil {
		panic("No \"vk_token\" field provided in config")
	}

	if config.VKChatId == nil {
		panic("No \"vk_chat_id\" field provided in config")
	}

	if config.Time == nil {
		panic("No \"time\" field provided in config")
	}

	if len(config.Servers) == 0 {
		panic("No \"servers\" field provided in config")
	}

	servers := make([]*ConfigServer, len(config.Servers))

	for id, server := range config.Servers {
		if server.Addr == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"addr\" field")
		}

		if server.Protocol == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"protocol\" field")
		}

		if server.Name == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"name\" field")
		}

		if server.MentionsText == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"mentions_text\" field")
		}

		if *server.Protocol != "udp" && *server.Protocol != "tcp" {
			panic("Server #" + strconv.Itoa(id) + " has wrong \"protocol\" field. It can be only \"udp\" or \"tcp\"")
		}

		if *server.Protocol == "udp" && server.StartBytesBase64 == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"start_bytes_base64\" field")
		}

		servers[id] = &ConfigServer{
			Name:             *server.Name,
			Addr:             *server.Addr,
			Protocol:         *server.Protocol,
			MentionsText:     *server.MentionsText,
			StartBytesBase64: *server.StartBytesBase64,
		}
	}

	return &ConfigStruct{
		VKToken:  *config.VKToken,
		Time:     *config.Time,
		VKChatId: *config.VKChatId,
		Servers:  servers,
	}
}
