package main

import (
	"github.com/cloudhonk/faras/khel"
	"github.com/cloudhonk/faras/renderer"
	"github.com/cloudhonk/faras/server"
)

func main() {

	farasInstance := khel.NewFaras()

	farasFrameConfig := renderer.FarasFrameConfig{
		Width:   80,
		Height:  25,
		Padding: 2,
	}
	farasPlayerManager := khel.NewFarasPlayerManager()
	farasFrameBuilder := renderer.NewFarasFrameBuilder(farasFrameConfig, farasPlayerManager)
	farasGameManager := server.NewFarasGameManager(farasFrameBuilder, farasInstance, farasPlayerManager)
	gameServer := server.NewGameServer(farasGameManager)

	go gameServer.StartServer()

	select {} // Keep the main function running
}
