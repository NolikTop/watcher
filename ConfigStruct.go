package main

// Существует только для проверки на наличие required полей в конфиге
type RawConfigStruct struct {
	VKToken  *string            `json:"vk_token"`
	VKChatId *int               `json:"vk_chat_id"`
	Servers  []*RawConfigServer `json:"servers"`
}

type ConfigStruct struct {
	VKToken  string          `json:"vk_token"`
	VKChatId int             `json:"vk_chat_id"`
	Servers  []*ConfigServer `json:"servers"`
}

// Существует только для проверки на наличие required полей в конфиге
type RawConfigServer struct {
	Addr     *string `json:"addr"`
	Protocol *string `json:"protocol"`
}

type ConfigServer struct {
	Addr     string `json:"addr"`
	Protocol string `json:"protocol"`
}
