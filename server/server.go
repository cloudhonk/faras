package server

import (
	"fmt"
	"net"

	"github.com/cloudhonk/faras/game"
	"github.com/cloudhonk/faras/logger"
)

type GameManager interface {
	Begin(juwadeyChan chan *game.Juwadey)
}

type GameServer struct {
	Manager     GameManager
	juwadeyChan chan *game.Juwadey
}

func NewGameServer(manager GameManager) *GameServer {
	s := GameServer{
		Manager:     manager,
		juwadeyChan: make(chan *game.Juwadey),
	}

	return &s
}

func (s *GameServer) StartServer() error {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		logger.Log.Error(fmt.Sprintf("error statring server: %s", err))
		return err
	}

	defer func() {
		if err := listener.Close(); err != nil {
			logger.Log.Error(fmt.Sprintf("error closing listener: %s", err))
		}
	}()

	go s.Manager.Begin(s.juwadeyChan)
	logger.Log.Info("Server started. Waiting for players...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Log.Error(fmt.Sprintf("error accepting connection: %s", err))
			continue
		}

		go s.handleConnection(conn)

	}
}

func (s *GameServer) handleConnection(conn net.Conn) {

	var name string
	conn.Write([]byte("Enter your name: "))
	fmt.Fscanln(conn, &name)
	juwadey := game.NewJuwadey(name, conn)

	s.juwadeyChan <- juwadey

}
