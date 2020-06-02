package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	configPtr := flag.String("config", "no", "path to JSON config")
	flag.Parse()
	config := parseConfig(*configPtr)

	fmt.Println(config)
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

	err = json.Unmarshal([]byte(file), config)
	if err != nil {
		panic("Couldn't read config file: " + err.Error())
	}

	if config.VKToken == nil {
		panic("No \"vk_token\" field provided in config")
	}

	if config.VKChatId == nil {
		panic("No \"vk_chat_id\" field provided in config")
	}

	if len(config.Servers) == 0 {
		panic("No \"servers\" field provided in config")
	}

	servers := make([]*ConfigServer, len(config.Servers))

	for id, server := range config.Servers {
		if server.Addr == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"Addr\" field")
		}

		if server.Protocol == nil {
			panic("Server #" + strconv.Itoa(id) + " hasn't \"Protocol\" field")
		}

		if *server.Protocol != "udp" && *server.Protocol != "tcp" {
			panic("Server #" + strconv.Itoa(id) + " has wrong \"Protocol\" field. It can be only \"udp\" or \"tcp\"")
		}

		servers[id] = &ConfigServer{
			Addr:     *server.Addr,
			Protocol: *server.Protocol,
		}
	}

	return &ConfigStruct{
		VKToken:  *config.VKToken,
		VKChatId: *config.VKChatId,
		Servers:  servers,
	}
}
