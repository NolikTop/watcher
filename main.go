package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"io/ioutil"
	"strconv"
)

func main() {
	configPtr := flag.String("config", "no", "path to JSON config")
	flag.Parse()
	config := parseConfig(*configPtr)

	initVK(config.VKToken, config.VKChatId)
	var typedServer Server

	for _, server := range config.Servers {
		if server.Protocol == "udp" {
			startBytes, err := base64.StdEncoding.DecodeString(server.StartBytesBase64)
			if err != nil {
				panic("Couldn't decode \"start_bytes_base64\" from server \"" + server.Name + "\"")
			}

			typedServer = &UDPServer{&ServerBase{}, startBytes, nil, nil}
		} else {
			typedServer = &TCPServer{&ServerBase{}}
		}

		typedServer.init(server.Name, server.Addr, server.MentionsText)
		go watch(typedServer, config.Time)
	}

	select {}
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
