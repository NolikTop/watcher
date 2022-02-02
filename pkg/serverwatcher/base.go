package serverwatcher

import (
	"math/rand"
	"watcher/pkg/notification"
)

type WatcherBase struct {
	name         string
	serverAddr   string
	working      bool
	offTime      uint
	mentionsText string
}

func (watcher *WatcherBase) SetWorking(working bool) {
	watcher.offTime = 0
	watcher.working = working
}

func (watcher *WatcherBase) IsWorking() bool {
	return watcher.working
}

func (watcher *WatcherBase) GetName() string {
	return watcher.name
}

func (watcher *WatcherBase) GetAddr() string {
	return watcher.serverAddr
}

func (watcher *WatcherBase) GetMentionsText() string {
	return watcher.mentionsText
}

func (watcher *WatcherBase) IncrementOffTime() {
	watcher.offTime++

	// todo внизу происходит что-то странное, надо здесь больше ясности внести
	if watcher.offTime&0b111 == 0b111 && (watcher.offTime < 30 || rand.Intn(4) == 1) {
		notification.SendBadNotification(watcher.name, watcher.serverAddr, watcher.mentionsText, watcher.offTime)
	}
}
