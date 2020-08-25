package main

type Server interface {
	init(name, addr, mentionsText string)
	getName() string
	getAddr() string
	checkConnection() error
	setWorking(working bool)
	isWorking() bool
	incrementOffTime()
	getMentionsText() string
}
