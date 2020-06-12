package main

// Существует только для проверки на наличие required полей в конфиге
type RawConfigStruct struct {
	VKToken  *string            `json:"vk_token"`
	VKChatId *int               `json:"vk_chat_id"`
	Time     *int               `json:"time"`
	Servers  []*RawConfigServer `json:"servers"`
}

type ConfigStruct struct {
	VKToken  string
	VKChatId int
	Time     int
	Servers  []*ConfigServer
}

// Существует только для проверки на наличие required полей в конфиге
type RawConfigServer struct {
	Name             *string `json:"name"`
	Addr             *string `json:"addr"`
	Protocol         *string `json:"protocol"`
	MentionsText     *string `json:"mentions_text"`
	StartBytesBase64 *string `json:"start_bytes_base64"`
}

type ConfigServer struct {
	Name             string
	Addr             string
	Protocol         string
	MentionsText     string
	StartBytesBase64 string
}
