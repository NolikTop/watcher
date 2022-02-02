package notification

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"watcher/pkg/serverwatcher"
)

type Vk struct {
	name string

	client *http.Client

	chatId      int
	accessToken string
}

func (v *Vk) Init(name string, data map[string]interface{}) error {
	v.name = name

	if chatId, ok := data["chat_id"]; ok {
		v.chatId = chatId.(int)
	} else {
		return errNoFieldInData("chat_id")
	}

	if accessToken, ok := data["access_token"]; ok {
		v.accessToken = accessToken.(string)
	} else {
		return errNoFieldInData("access_token")
	}

	v.client = &http.Client{}

	return nil
}

func (v *Vk) GetName() string {
	return v.name
}

func (v *Vk) NotifyServerWentDown(server serverwatcher.ServerWatcher) error {
	message := fmt.Sprintf(
		`Сервер %s упал.
Призываю %s`,
		server.GetFormattedName(), server.GetMentionsText(),
	)

	return v.sendMessage(message)
}

func (v *Vk) NotifyServerStillIsDown(server serverwatcher.ServerWatcher) error {
	message := fmt.Sprintf(
		`Сервер %s все еще лежит. Прошло уже %d сек.
Призываю %s`,
		server.GetName(), server.GetOffTime(), server.GetMentionsText(),
	)

	return v.sendMessage(message)
}

func (v *Vk) NotifyServerIsUp(server serverwatcher.ServerWatcher) error {
	message := fmt.Sprintf(
		`Сервер %s встал.
Призываю %s`,
		server.GetName(), server.GetMentionsText(),
	)

	return v.sendMessage(message)
}

func (v *Vk) sendMessage(message string) error {
	request, err := http.NewRequest("GET", "https://api.vk.com/method/messages.send", nil)
	if err != nil {
		return err
	}

	query := request.URL.Query()
	query.Add("v", "5.107") // todo вынести указание версии API в конфиг
	query.Add("chat_id", strconv.Itoa(v.chatId))
	query.Add("access_token", v.accessToken)
	query.Add("random_id", "0")
	query.Add("message", message)
	request.URL.RawQuery = query.Encode()

	response, err := v.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = getErrorFromResponse(body)
	if err != nil {
		return err
	}

	return nil
}

func getErrorFromResponse(responseBody []byte) error {
	//todo
	return nil
}
