package server

import (
	"fmt"
	"net"
)

const (
	MAX_PLAYERS = 4
)

type GameManager interface {
	Start(conn net.Conn)
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
		go s.Manager.Start(conn)

	}
}

// func (s *GameServer) handleConnection(conn net.Conn) {
// 	if condition := s.addPlayer(conn); !condition {
// 		conn.Close()
// 		return
// 	}
// 	if len(s.GameInstance.Juwadeys) == MAX_PLAYERS {
// 		go s.GameInstance.GameLoop()
// 	}

// 	// Block until we receive a signal that the game has ended
// 	select {
// 	case <-s.GameInstance.End:
// 		s.GameInstance.Reset()
// 		return
// 	}
// }

// func (s *GameServer) addPlayer(conn net.Conn) bool {

// 	var playerName string
// 	fmt.Fprintln(conn, "Enter your name: ")
// 	fmt.Fscanln(conn, &playerName)

// 	s.mu.Lock()
// 	if len(s.GameInstance.Juwadeys) >= MAX_PLAYERS {
// 		fmt.Fprintln(conn, "The game is full!")
// 		s.mu.Unlock()
// 		return false
// 	}

// 	juwadey := khel.NewJuwadey(playerName, conn)
// 	s.GameInstance.Juwadeys = append(s.GameInstance.Juwadeys, juwadey)
// 	s.mu.Unlock()

// 	fmt.Fprintf(conn, "Welcome, %s! Waiting for other players...\n", playerName)
// 	return true
// }
