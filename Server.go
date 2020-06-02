package main

type Server interface {
	checkConnection() error
}
