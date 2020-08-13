package main

import (
	"encoding/base64"
	"flag"
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
