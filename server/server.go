package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/cloudhonk/faras/khel"
)

const (
	MAX_PLAYERS = 4
)

type GameServer struct {
	GameInstance *khel.Faras
	mu           sync.Mutex
}

func NewGameServer(instance *khel.Faras) *GameServer {
	s := GameServer{
		GameInstance: instance,
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

		go s.handleConnection(conn)
	}
}

func (s *GameServer) handleConnection(conn net.Conn) {
<<<<<<< HEAD
	defer conn.Close()
	s.addPlayer(conn)
=======
	if condition := s.addPlayer(conn); !condition {
		conn.Close()
		return
	}
>>>>>>> 55595c09f67facc42903b4ee746dc865d38065d0
	if len(s.GameInstance.Juwadeys) == MAX_PLAYERS {
		go s.GameInstance.GameLoop()
	}

<<<<<<< HEAD
	// Keep the connection open
	select {}
}

func (s *GameServer) addPlayer(conn net.Conn) {
=======
	// Block until we receive a signal that the game has ended
	select {
	case <-s.GameInstance.End:
		s.GameInstance.Reset()
		return
	}
}

func (s *GameServer) addPlayer(conn net.Conn) bool {
>>>>>>> 55595c09f67facc42903b4ee746dc865d38065d0

	var playerName string
	fmt.Fprintln(conn, "Enter your name: ")
	fmt.Fscanln(conn, &playerName)

	s.mu.Lock()
	if len(s.GameInstance.Juwadeys) >= MAX_PLAYERS {
		fmt.Fprintln(conn, "The game is full!")
		s.mu.Unlock()
<<<<<<< HEAD
		return
=======
		return false
>>>>>>> 55595c09f67facc42903b4ee746dc865d38065d0
	}

	juwadey := khel.NewJuwadey(playerName, conn)
	s.GameInstance.Juwadeys = append(s.GameInstance.Juwadeys, juwadey)
	s.mu.Unlock()

	fmt.Fprintf(conn, "Welcome, %s! Waiting for other players...\n", playerName)
<<<<<<< HEAD

=======
	return true
>>>>>>> 55595c09f67facc42903b4ee746dc865d38065d0
}
