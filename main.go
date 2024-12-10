package main

import (
	"github.com/cloudhonk/faras/game"
	"github.com/cloudhonk/faras/server"
)

func main() {

	farasGameManager := game.NewFarasGameManager()
	gameServer := server.NewGameServer(farasGameManager)

	err := gameServer.StartServer()

	if err != nil {
		panic("Failed to start server")
	}

	select {}
}
