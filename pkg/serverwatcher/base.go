package serverwatcher

import "fmt"

type WatcherBase struct {
	name         string
	serverAddr   string
	protocol     string
	chats        []string
	working      bool
	offTime      uint
	mentionsText string
}

func (w *WatcherBase) SetWorking(working bool) {
	w.offTime = 0
	w.working = working
}

func (w *WatcherBase) IsWorking() bool {
	return w.working
}

func (w *WatcherBase) GetName() string {
	return w.name
}

func (w *WatcherBase) GetServerAddr() string {
	return w.serverAddr
}

func (w *WatcherBase) GetProtocol() string {
	return w.protocol
}

func (w *WatcherBase) GetChats() []string {
	return w.chats
}

func (w *WatcherBase) GetTimeout() int {
	return 10 // todo
}

func (w *WatcherBase) GetFormattedName() string {
	return fmt.Sprintf("%s (%s)", w.GetName(), w.GetProtocol())
}

func (w *WatcherBase) GetMentionsText() string {
	return w.mentionsText
}

func (w *WatcherBase) IncrementOffTime() {
	w.offTime++
}

func (w *WatcherBase) GetOffTime() uint {
	return w.offTime
}
