package serverwatcher

type ServerWatcher interface {
	Init(data map[string]interface{}) error
	CheckConnection() error

	SetWorking(working bool)
	IsWorking() bool
	GetName() string
	GetAddr() string
	GetMentionsText() string
	IncrementOffTime()
}
