package chat

import (
	"fmt"
	"github.com/NolikTop/watcher/pkg/server"
	"io"
	"net/http"
	"strconv"
)

type Tg struct {
	name string

	client *http.Client

	chatId      int
	accessToken string
}

func (v *Tg) Init(name string, data map[string]interface{}) error {
	v.name = name

	if chatId, ok := data["chat_id"]; ok {
		v.chatId = int(chatId.(float64)) // сразу в int не дает кастить =(((
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

func (v *Tg) GetName() string {
	return v.name
}

func (v *Tg) NotifyServerWentDown(server server.Server, err error) error {
	message := fmt.Sprintf(
		`Сервер %s упал. 
Причина: %s
Призываю %s`,
		server.GetFormattedName(), err.Error(), server.GetMentionsText(),
	)

	return v.sendMessage(message)
}

func (v *Tg) NotifyServerStillIsDown(server server.Server) error {
	message := fmt.Sprintf(
		`Сервер %s все еще лежит. Прошло уже %d сек.
Призываю %s`,
		server.GetFormattedName(), server.GetOffTime(), server.GetMentionsText(),
	)

	return v.sendMessage(message)
}

func (v *Tg) NotifyServerIsUp(server server.Server) error {
	message := fmt.Sprintf(
		`Сервер %s встал.
Призываю %s`,
		server.GetFormattedName(), server.GetMentionsText(),
	)

	return v.sendMessage(message)
}

func (v *Tg) sendMessage(message string) error {
	request, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", v.accessToken), nil)
	if err != nil {
		return err
	}

	query := request.URL.Query()
	query.Add("chat_id", strconv.Itoa(v.chatId))
	query.Add("text", message)
	request.URL.RawQuery = query.Encode()

	response, err := v.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = v.getErrorFromResponse(body)
	if err != nil {
		return err
	}

	return nil
}

func (v *Tg) getErrorFromResponse(responseBody []byte) error {
	//todo
	return nil
}
