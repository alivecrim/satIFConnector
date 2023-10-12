package main

import "pasha.com/satifconnector/internal/app/server"

func main() {
	newServer := server.NewServer()
	newServer.Init()
}
