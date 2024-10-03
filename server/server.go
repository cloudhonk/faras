package server

import (
	"fmt"
	"net"
)

const (
	MAX_PLAYERS = 4
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

func (s *GameServer) StartServer() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for players...")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go s.Manager.Join(conn)

	}
}
