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
	Command          *string `json:"command"`
	Password         *string `json:"password"`
}

type ConfigServer struct {
	Name             string
	Addr             string
	Protocol         string
	MentionsText     string
	StartBytesBase64 string
	Command          string
	Password         string
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

	var server *ConfigServer

	for id, iteratingServer := range config.Servers {
		if iteratingServer.Addr == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"addr\" field")
		}

		if iteratingServer.Protocol == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"protocol\" field")
		}

		if iteratingServer.Name == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"name\" field")
		}

		if iteratingServer.MentionsText == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"mentions_text\" field")
		}

		switch *iteratingServer.Protocol {
		case "udp":
			if iteratingServer.StartBytesBase64 == nil {
				panic("Server #" + strconv.Itoa(id) + " hasn't \"start_bytes_base64\" field")
			}

			server = &ConfigServer{
				Name:             *iteratingServer.Name,
				Addr:             *iteratingServer.Addr,
				Protocol:         *iteratingServer.Protocol,
				MentionsText:     *iteratingServer.MentionsText,
				StartBytesBase64: *iteratingServer.StartBytesBase64,
			}
		case "tcp":
			fallthrough
		case "minecraft":
			server = &ConfigServer{
				Name:         *iteratingServer.Name,
				Addr:         *iteratingServer.Addr,
				Protocol:     *iteratingServer.Protocol,
				MentionsText: *iteratingServer.MentionsText,
			}
		case "rcon":
			server = &ConfigServer{
				Name:         *iteratingServer.Name,
				Addr:         *iteratingServer.Addr,
				Protocol:     *iteratingServer.Protocol,
				MentionsText: *iteratingServer.MentionsText,
				Command:      *iteratingServer.Command,
				Password:     *iteratingServer.Password,
			}
		default:
			panic("Server #" + strconv.Itoa(id) + " has wrong \"protocol\" field. It can be only \"udp\", \"tcp\", \"minecraft\" or \"rcon\"")
		}

		servers[id] = server
	}

	return &ConfigStruct{
		VKToken:  *config.VKToken,
		Time:     *config.Time,
		VKChatId: *config.VKChatId,
		Servers:  servers,
	}
}
