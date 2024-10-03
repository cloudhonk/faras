package main

import (
	"github.com/cloudhonk/faras/game"
	"github.com/cloudhonk/faras/renderer"
	"github.com/cloudhonk/faras/server"
)

func main() {

	farasFrameConfig := renderer.FarasFrameConfig{
		Width:   renderer.WINDOW_WIDTH,
		Height:  renderer.WINDOW_HEIGHT,
		Padding: renderer.WINDOW_PADDING,
	}

	farasFrameBuilder := renderer.NewFarasFrameBuilder(farasFrameConfig)
	farasGameManager := game.NewFarasGameManager(farasFrameBuilder)
	gameServer := server.NewGameServer(farasGameManager)

	go gameServer.StartServer()
	go farasGameManager.Update()
	go farasGameManager.End()

	select {}
}
