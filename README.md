## watcher

Watcher следит за указанными ему tcp/udp серверами, и если какие-то из серверов прекратили работу, watcher предупредит Вас об этом в ВКонтакте

![watcher works](https://sun1-95.userapi.com/EeKVmoN8KkpstL0xCDJ0iHr68BjGaOVoFSvI1Q/rRLlzTQUsOA.jpg)

---

## Установка

Для нормальной работы Watcher рекомендуется иметь версию Go 1.14 или выше

```shell script
go get github.com/GreenWix/watch
cd $GOPATH/src/github.com/GreenWix/watch
go build -o watch
```

## Запуск

```shell script
./watch -config=путь/до/конфига.json
```

---

## Конфиг

Поле | Описание
------------ | -------------
vk_token | токен группы ВКонтакте
vk_chat_id | id чата ВКонтакте, в который будут присылаться уведомления
time | Раз в это время (в секундах) будет осуществляться проверка серверов (при условии, что последняя проверка watcher'a была успешной)
servers | сервера

Описание полей сервера

Поле | Описание
------------ | -------------
name | имя сервера (нужно для уведомлений, чтобы Вы поняли какой именно сервер упал)
addr | адрес сервера 
protocol | tcp/udp/minecraft. Протокол, используемый сервером
mentions_text | строка, содержащая список упоминаний пользователей, которые ответственны за данный сервер
start_bytes_base64 (только UDP) | Base64 байтов, после отправки которых watcher должен получить ответ от UDP сервера

Если последняя проверка watcher'а была неуспешной, то watcher будет проверять каждую секунду, пока результат проверки не станет успешным

Пример конфига:
```json
{
  "vk_token": "token",
  "vk_chat_id": 1,
  "time": 10,
  "servers": [
    {
      "name": "my minecraft pocket edition or bedrock edition server",
      "addr": "127.0.0.1:19132",
      "protocol": "minecraft",
      "mentions_text": "@online"
    },
    {
      "name": "my http server",
      "addr": "127.0.0.1:8080",
      "protocol": "tcp",
      "mentions_text": "@all"
    },
    {
      "name": "my udp server",
      "addr": "127.0.0.1:19132",
      "protocol": "udp",
      "start_bytes_base64": "BQD//wD+/v7+/f39/RI0VngKAA==",
      "mentions_text": "@online"
    }
  ]
}
```
