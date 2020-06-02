## watcher

Watcher следит за указанными ему tcp/udp серверами, и если какие-то из серверов прекратили работу, watcher предупредит Вас об этом в ВКонтакте

## Установка

```shell script
go get github.com/GreenWix/watcher
go build -o main
```

## Запуск

```shell script
./main config=/абсолютный/путь/до/конфига.json
```

## TODO список
+ Поддержка UDP серверов
+ Указание байтов, после отправки которых UDP серверам, должен быть получен ответ

## Примеры конфига

```json
{
  "vk_token": "token",
  "vk_chat_id": 1,
  "time": 10,
  "servers": [
    {
      "name": "my http server",
      "addr": "127.0.0.1:8080",
      "protocol": "tcp",
      "mentions_text": "@all"
    },
    {
      "name": "my udp server",
      "addr": "127.0.0.1:1234",
      "protocol": "udp",
      "mentions_text": "@online"
    }
  ]
}
```