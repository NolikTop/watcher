package watcher

import (
	"github.com/NolikTop/watcher/pkg/chat"
	"github.com/NolikTop/watcher/pkg/config"
	"github.com/NolikTop/watcher/pkg/notification"
	"github.com/NolikTop/watcher/pkg/server"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync/atomic"
	"time"
)

type Watcher struct {
	servers []server.Server
	chats   map[string]chat.Chat

	canBeChanged *atomic.Value
}

func New() *Watcher {
	val := &atomic.Value{}
	val.Store(true)

	return &Watcher{
		servers:      []server.Server{},
		chats:        make(map[string]chat.Chat),
		canBeChanged: val,
	}
}

func (w *Watcher) Load(config *config.Config) error {
	log.Info("Loading notification chats...")

	err := w.addChatsFromConfig(config.Chats)
	if err != nil {
		return err
	}

	log.Info("Loading servers...")

	err = w.addServersFromConfig(config.Servers)
	if err != nil {
		return err
	}

	return nil
}

func (w *Watcher) addChatsFromConfig(chatConfigs []*config.ChatConfig) error {
	for _, chatConfig := range chatConfigs {
		cht, err := chat.NewChat(chatConfig)
		if err != nil {
			return err
		}

		err = w.AddChat(cht)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *Watcher) addServersFromConfig(serversConfigs []*config.ServerConfig) error {
	for _, serverConfig := range serversConfigs {
		serv, err := server.NewServer(serverConfig)
		if err != nil {
			return err
		}

		err = w.AddServer(serv)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *Watcher) GetServer() []server.Server {
	return w.servers
}

func (w *Watcher) GetChats() map[string]chat.Chat {
	return w.chats
}

func (w *Watcher) GetChat(name string) chat.Chat {
	return w.chats[name]
}

// AddServer не является потокобезопасным (да и вообще кому в голову может прийти добавлять сервера в разных потоках?)
func (w *Watcher) AddServer(serv server.Server) error {
	if err := w.checkCanBeChanged(); err != nil {
		return err
	}

	w.servers = append(w.servers, serv)
	//todo

	return nil
}

// AddChat не является потокобезопасным (да и вообще кому в голову может прийти добавлять методы в разных потоках?)
func (w *Watcher) AddChat(cht chat.Chat) error {
	if err := w.checkCanBeChanged(); err != nil {
		return err
	}

	name := cht.GetName()
	if _, ok := w.chats[name]; ok {
		return errChatWithThisNameAlreadyExists(name)
	}

	w.chats[name] = cht

	return nil
}

func (w *Watcher) checkCanBeChanged() error {
	if w.canBeChanged.Load().(bool) {
		return nil
	}

	return errWatcherCantBeChanged
}

func (w *Watcher) Start() error {
	w.canBeChanged.Store(false)

	err := w.checkChatNamesInServers()
	if err != nil {
		return err
	}

	w.runAllServerWatchers()

	return nil
}

func (w *Watcher) checkChatNamesInServers() error {
	chats := w.GetChats()

	for _, serv := range w.servers {
		for _, chatName := range serv.GetChatNames() {
			if _, ok := chats[chatName]; !ok {
				return errUnknownChatName(serv.GetName(), chatName)
			}

		}
	}

	return nil
}

func (w *Watcher) runAllServerWatchers() {
	for _, serv := range w.servers {
		go w.Watch(serv)
	}
}

func (w *Watcher) Watch(serv server.Server) {
	var err error
	timeout := serv.GetTimeout()

	for {
		err = serv.CheckConnection()

		if serv.IsMarkedAsWorking() {
			if err != nil {
				w.serverWentDown(serv, err)
			}
			time.Sleep(time.Duration(timeout) * time.Second)
		} else {
			if err == nil {
				w.serverStartedUp(serv)
			} else {
				serv.IncrementOffTime()
				if w.shouldReportAgain(serv) {
					notification.ServerStillIsDown(w, serv)
				}
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func (w *Watcher) shouldReportAgain(server server.Server) bool {
	// todo эту уникальную формулу явно стоит переделать
	return server.GetOffTime()&0b111 == 0b111 && (server.GetOffTime() < 30 || rand.Intn(4) == 1)
}

func (w *Watcher) serverWentDown(server server.Server, err error) {
	server.MarkIsWorking(false)
	log.WithField("server", server.GetFormattedName()).WithField("error", err).Info("Server went down")
	notification.ServerWentDown(w, server, err)
}

func (w *Watcher) serverStartedUp(server server.Server) {
	server.MarkIsWorking(true)
	log.WithField("server", server.GetFormattedName()).Info("Server is working again")
	notification.ServerIsUp(w, server)
}
