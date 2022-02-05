package server

type Server interface {
	Init(data map[string]interface{}) error
	CheckConnection() error

	GetName() string
	GetServerAddr() string
	GetProtocol() string

	GetChats() []string

	GetTimeout() int

	GetFormattedName() string

	GetMentionsText() string

	MarkIsWorking(working bool)
	IsMarkedAsWorking() bool

	IncrementOffTime()
	GetOffTime() uint
}
