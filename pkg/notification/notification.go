package notification

import (
	"github.com/NolikTop/watcher/pkg/server"
	log "github.com/sirupsen/logrus"
)

func ServerWentDown(chats chatsContainer, serv server.Server, connectionErr error) {
	for _, chatName := range serv.GetChatNames() {
		cht := chats.GetChat(chatName)

		err := cht.NotifyServerWentDown(serv, connectionErr)
		if err != nil {
			log.Error(err)
		}
	}
}

func ServerStillIsDown(chats chatsContainer, serv server.Server) {
	for _, chatName := range serv.GetChatNames() {
		cht := chats.GetChat(chatName)

		err := cht.NotifyServerStillIsDown(serv)
		if err != nil {
			log.Error(err)
		}
	}
}

func ServerIsUp(chats chatsContainer, serv server.Server) {
	for _, chatName := range serv.GetChatNames() {
		cht := chats.GetChat(chatName)

		err := cht.NotifyServerIsUp(serv)
		if err != nil {
			log.Error(err)
		}
	}
}
