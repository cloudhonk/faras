package server

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/cloudhonk/faras/bung"
	"github.com/cloudhonk/faras/khel"
)

const (
	MAX_ROUNDS = 2
)

type FrameBuilder interface {
	Build(playerRef int)
	Flush() []byte
}

type PlayerManager interface {
	AddPlayer(*khel.Juwadey)
	GetPlayers() []*khel.Juwadey
	CheckIfFull() bool
	DealCardToPlayer(card *bung.Taas, playerIdx int)
}

type FarasGameManager struct {
	FrameBuilder
	PlayerManager
	GameInstance *khel.Faras
	mu           sync.Mutex
}

func NewFarasGameManager(fb FrameBuilder, gameInstance *khel.Faras, playerManager PlayerManager) *FarasGameManager {
	return &FarasGameManager{
		GameInstance:  gameInstance,
		FrameBuilder:  fb,
		PlayerManager: playerManager,
	}
}

// func (fgm *FarasGameManager) AddPlayer(conn net.Conn) bool {
// 	var playerName string
// 	conn.Write([]byte("Enter your name: "))
// 	fmt.Fscanln(conn, &playerName)

// 	fgm.mu.Lock()
// 	if len(fgm.GameInstance.Juwadeys) >= MAX_PLAYERS {
// 		conn.Write([]byte("The game is full!"))
// 		fgm.mu.Unlock()
// 		return false
// 	}

// 	juwadey := khel.NewJuwadey(playerName, conn)
// 	fgm.GameInstance.Juwadeys = append(fgm.GameInstance.Juwadeys, juwadey)
// 	fgm.mu.Unlock()

// 	fgm.FrameBuilder.Build()
// 	return true
// }

func (fgm *FarasGameManager) Start(conn net.Conn) {

	fgm.mu.Lock()
	defer fgm.mu.Unlock()
	full := fgm.CheckIfFull()
	if full {
		conn.Write([]byte("The game is full!"))
		conn.Close()
		return
	}

	var playerName string
	conn.Write([]byte("Enter your name: "))
	fmt.Fscanln(conn, &playerName)

	juwadey := khel.NewJuwadey(playerName, conn)
	// juwadey.Haat = []*bung.Taas{{Patti: bung.Attha, Rangi: bung.Hukum}, {Patti: bung.Nahar, Rangi: bung.Hukum}, {Patti: bung.Ekka, Rangi: bung.Itta}}
	fgm.PlayerManager.AddPlayer(juwadey)
	fgm.Broadcast()

	if len(fgm.PlayerManager.GetPlayers()) == MAX_PLAYERS {
		go fgm.BeginGame()
	}
}

func (fgm *FarasGameManager) BeginGame() {
	deck := bung.New()
	for i := range 3 {
		for j, _ := range fgm.PlayerManager.GetPlayers() {
			fgm.PlayerManager.DealCardToPlayer(&deck[j+3*i], j)
			time.Sleep(3 * time.Second)
			fgm.Broadcast()
		}
	}
}

func (fgm *FarasGameManager) Broadcast() {
	for i, juwadey := range fgm.PlayerManager.GetPlayers() {
		fgm.FrameBuilder.Build(i)
		juwadey.Conn.Write(fgm.Flush())
	}
}

func (fgm *FarasGameManager) Update() {

}

func (fgm *FarasGameManager) End() {

}
