package server

import (
	"fmt"
	"net"

	"github.com/cloudhonk/faras/logger"
)

type GameManager interface {
	Join(conn net.Conn)
	Update()
	End()
}

type GameServer struct {
	Manager GameManager
}

func NewGameServer(manager GameManager) *GameServer {
	s := GameServer{
		Manager: manager,
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

	go s.Manager.Update()
	go s.Manager.End()

	logger.Log.Info("Server started. Waiting for players...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Log.Error(fmt.Sprintf("error accepting connection: %s", err))
			continue
		}

		go s.Manager.Join(conn)

	}
}
