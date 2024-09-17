package main

import (
	"github.com/cloudhonk/faras/khel"
	"github.com/cloudhonk/faras/server"
)

func main() {
	farasInstance := khel.NewFaras()
	gameServer := server.NewGameServer(farasInstance)

	go gameServer.StartServer()

	select {} // Keep the main function running
}
