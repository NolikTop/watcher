package config

// RawConfig существует только для проверки на наличие required полей в конфиге
type RawConfig struct {
	NotificationMethods []*RawNotificationMethodConfig `json:"notification_methods"`
	Servers             []*RawServerConfig             `json:"servers"`
}

type Config struct {
	NotificationMethods []*NotificationMethodConfig
	Servers             []*ServerConfig
}

// RawNotificationMethodConfig существует только для проверки на наличие required полей в конфиге
type RawNotificationMethodConfig struct {
	Name   *string                 `json:"name"`
	Method *string                 `json:"method"`
	Data   *map[string]interface{} `json:"data"`
}

type NotificationMethodConfig struct {
	Name   string
	Method string
	Data   map[string]interface{}
}

// RawServerConfig существует только для проверки на наличие required полей в конфиге
type RawServerConfig struct {
	Name         *string                 `json:"name"`
	Addr         *string                 `json:"addr"`
	Protocol     *string                 `json:"protocol"`
	Chats        *[]string               `json:"chats"`
	MentionsText *string                 `json:"mentions_text"`
	Data         *map[string]interface{} `json:"data"`
}

type ServerConfig struct {
	Name         string
	Addr         string
	Protocol     string
	Chats        []string
	MentionsText string
	Data         map[string]interface{}
}
