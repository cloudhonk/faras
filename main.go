package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

type Bung []Taas

func New() Bung {
	bung := Bung{}
	for _, patti := range Pattiharu {
		for _, turup := range Turupharu {
			bung = append(bung, Taas{Patti: patti, Turup: turup})
		}
	}
	bung.Fitt()
	return bung
}

func (b *Bung) Fitt() {
	rand.Shuffle(len(*b), func(i, j int) { (*b)[i], (*b)[j] = (*b)[j], (*b)[i] })
}

type Patti int

const (
	Dua = iota + 2
	Tirki
	Chauka
	Panja
	Chhaka
	Satta
	Attha
	Nahar
	Dahar
	Gulam
	Missi
	Bassa
	Ekka
)

var Pattiharu = []Patti{Dua, Tirki, Chauka, Panja, Chhaka, Satta, Attha, Nahar, Dahar, Gulam, Missi, Bassa, Ekka}

func (p *Patti) String() string {
	switch *p {

	case Gulam:
		return "J"

	case Missi:
		return "Q"

	case Bassa:
		return "K"

	case Ekka:
		return "A"
	}
	return fmt.Sprintf("%d", *p)
}

type Taas struct {
	Patti
	Turup
}

type Turup int

const (
	Paan = iota + 1
	Itta
	Chidi
	Hukum
)

var Turupharu = []Turup{Paan, Itta, Chidi, Hukum}

func (s *Turup) String() string {

	switch *s {

	case Paan:
		return "♥"

	case Itta:
		return "♦"

	case Chidi:
		return "♣"

	case Hukum:
		return "♠"
	}

	return ""

}

func (t *Taas) String() string {
	return fmt.Sprintf("%s को %s", t.Turup.String(), t.Patti.String())
}

// func main() {
// 	bung := New()

// 	for _, taas := range bung {
// 		fmt.Println("Hello", taas.String())
// 	}
// }

type Player struct {
	Name string
	Conn net.Conn
	Wins int
	Card Taas
}

type Server struct {
	Players []*Player
	Mutex   sync.Mutex
	Round   int
	Deck    Bung
}

const maxPlayers = 4
const totalRounds = 10

func (s *Server) StartServer() {
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

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	var playerName string
	fmt.Fprintln(conn, "Enter your name: ")
	fmt.Fscanln(conn, &playerName)

	s.Mutex.Lock()
	if len(s.Players) >= maxPlayers {
		fmt.Fprintln(conn, "The game is full!")
		s.Mutex.Unlock()
		return
	}

	player := &Player{Name: playerName, Conn: conn, Wins: 0}
	s.Players = append(s.Players, player)
	s.Mutex.Unlock()

	fmt.Fprintf(conn, "Welcome, %s! Waiting for other players...\n", playerName)

	if len(s.Players) == maxPlayers {
		go s.startGame()
	}

	// Keep the connection open
	select {}
}

func (s *Server) startGame() {
	for s.Round = 1; s.Round <= totalRounds; s.Round++ {
		s.dealCards()
		s.findWinner()
		s.updateScoreboard()
		time.Sleep(2 * time.Second)
	}

	s.endGame()
}

func (s *Server) dealCards() {
	s.Deck = New()

	for i, player := range s.Players {
		player.Card = s.Deck[i]
		fmt.Fprintf(player.Conn, "Round %d: You are dealt card number %s\n", s.Round, player.Card.String())
	}
}

func (s *Server) findWinner() {
	var winner *Player
	highestCard := -1

	for _, player := range s.Players {
		if int(player.Card.Patti) > highestCard {
			highestCard = int(player.Card.Patti)
			winner = player
		}
	}

	if winner != nil {
		winner.Wins++
		for _, player := range s.Players {
			fmt.Fprintf(player.Conn, "Round %d Winner: %s with card %s\n", s.Round, winner.Name, winner.Card.String())
		}
	}
}

func (s *Server) updateScoreboard() {
	for _, player := range s.Players {
		for _, p := range s.Players {
			fmt.Fprintf(p.Conn, "Scoreboard: %s has %d wins\n", player.Name, player.Wins)
		}
	}
}

func (s *Server) endGame() {
	var overallWinner *Player
	highestWins := -1

	for _, player := range s.Players {
		if player.Wins > highestWins {
			highestWins = player.Wins
			overallWinner = player
		}
	}

	for _, player := range s.Players {
		if overallWinner != nil {
			fmt.Fprintf(player.Conn, "The overall winner is %s with %d wins!\n", overallWinner.Name, overallWinner.Wins)
		}
	}

	fmt.Println("Game has ended. Server is shutting down.")
	os.Exit(0)
}

func main() {
	server := &Server{
		Players: []*Player{},
	}

	go server.StartServer()

	select {} // Keep the main function running
}
