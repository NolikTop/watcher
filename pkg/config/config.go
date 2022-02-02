package config

// RawConfig существует только для проверки на наличие required полей в конфиге
type RawConfig struct {
	VKToken  *string            `json:"vk_token"`
	VKChatId *int               `json:"vk_chat_id"`
	Time     *int               `json:"time"`
	Servers  []*RawServerConfig `json:"servers"`
}

type Config struct {
	VKToken  string
	VKChatId int
	Time     int
	Servers  []*ServerConfig
}

// RawServerConfig существует только для проверки на наличие required полей в конфиге
type RawServerConfig struct {
	Name         *string                 `json:"name"`
	Addr         *string                 `json:"addr"`
	Protocol     *string                 `json:"protocol"`
	MentionsText *string                 `json:"mentions_text"`
	Data         *map[string]interface{} `json:"data"`
}

type ServerConfig struct {
	Name         string
	Addr         string
	Protocol     string
	MentionsText string
	Data         map[string]interface{}
}
