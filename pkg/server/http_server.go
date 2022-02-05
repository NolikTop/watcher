package server

import (
	"net/http"
)

type HttpServer struct {
	*Base
}

func (s *HttpServer) Init(data map[string]interface{}) error {
	return nil
}

func (s *HttpServer) CheckConnection() (err error) {
	_, err = http.Get("http://" + s.serverAddr)
	if err != nil {
		return
	}

	return nil
}
