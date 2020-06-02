package main

import "math/rand"

type ServerBase struct {
	Name         string
	Addr         string
	Working      bool
	OffTime      uint
	MentionsText string
}

func (server *ServerBase) setWorking(working bool) {
	server.OffTime = 0
	server.Working = working
}

func (server *ServerBase) isWorking() bool {
	return server.Working
}

func (server *ServerBase) getName() string {
	return server.Name
}

func (server *ServerBase) getAddr() string {
	return server.Addr
}

func (server *ServerBase) getMentionsText() string {
	return server.MentionsText
}

func (server *ServerBase) init(name, addr, mentionsText string) {
	server.Name = name
	server.Addr = addr
	server.Working = true
	server.OffTime = 0
	server.MentionsText = mentionsText
}

func (server *ServerBase) incrementOffTime() {
	server.OffTime++
	if server.OffTime&0b111 == 0b111 && rand.Intn(4) == 1 {
		sendBadNotification(server.Name, server.Addr, server.MentionsText, server.OffTime)
	}
}
