package server

import "fmt"

type Base struct {
	name         string
	serverAddr   string
	protocol     string
	chats        []string
	working      bool
	offTime      uint
	mentionsText string
}

func (b *Base) SetWorking(working bool) {
	b.offTime = 0
	b.working = working
}

func (b *Base) IsWorking() bool {
	return b.working
}

func (b *Base) GetName() string {
	return b.name
}

func (b *Base) GetServerAddr() string {
	return b.serverAddr
}

func (b *Base) GetProtocol() string {
	return b.protocol
}

func (b *Base) GetChats() []string {
	return b.chats
}

func (b *Base) GetTimeout() int {
	return 10 // todo
}

func (b *Base) GetFormattedName() string {
	return fmt.Sprintf("%s (%s)", b.GetName(), b.GetProtocol())
}

func (b *Base) GetMentionsText() string {
	return b.mentionsText
}

func (b *Base) IncrementOffTime() {
	b.offTime++
}

func (b *Base) GetOffTime() uint {
	return b.offTime
}
