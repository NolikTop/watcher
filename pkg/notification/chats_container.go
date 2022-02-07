package notification

import "github.com/NolikTop/watcher/pkg/chat"

// не самое удачное решение, но пока я нашел только его
// нужно оно для решения проблемы циклического импорта
// дело в том, что мне нужно как-то получить список чатов, в которые
// будет писать о падении сервера, однако chat.Chat уже ссылается
// на server.Server (получает инфу о нем). Что с этим делать пока
// не очень понятно.
type chatsContainer interface {
	GetChat(name string) chat.Chat
}
