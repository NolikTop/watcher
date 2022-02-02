package notification

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var readyUrl = ""

func initVK(token string, chatId int) {
	readyUrl = "http://api.vk.com/method/messages.send?v=5.107&random_id=0&chat_id=" + strconv.Itoa(chatId) + "&access_token=" + url.QueryEscape(token) + "&message="
}

func SendErrorNotification(name string, addr string, mentionsText string, err error) {
	makeRequest("Сервер \"" + name + "\" (" + addr + ") упал.\nОшибка: " + err.Error() + "\nПризываю " + mentionsText)
}

func SendBadNotification(name string, addr string, mentionsText string, offTime uint) {
	if offTime < 100 { // чтобы только вначале упоминало
		makeRequest("Сервер \"" + name + "\" (" + addr + ") лежит уже " + strconv.Itoa(int(offTime)) + " секунд.\nПризываю " + mentionsText)
	} else {
		makeRequest("Сервер \"" + name + "\" (" + addr + ") лежит уже " + strconv.Itoa(int(offTime)) + " секунд")
	}
}

func SendOkNotification(name string, addr string) {
	makeRequest("Сервер \"" + name + "\" (" + addr + ") поднялся")
}

func makeRequest(ur string) {
	_, err := http.Get(readyUrl + url.QueryEscape(ur))
	if err != nil {
		fmt.Println("Couldn't make GET request: " + err.Error())
	}
}
