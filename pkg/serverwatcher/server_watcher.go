package serverwatcher

type ServerWatcher interface {
	Init(data map[string]interface{}) error
	CheckConnection() error

	GetName() string
	GetServerAddr() string
	GetProtocol() string

	GetChats() []string

	GetTimeout() int

	GetFormattedName() string

	GetMentionsText() string

	SetWorking(working bool)
	IsWorking() bool

	IncrementOffTime()
	GetOffTime() uint
}
