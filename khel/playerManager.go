package khel

import "github.com/cloudhonk/faras/bung"

type FarasPlayerManager struct {
	Players []*Juwadey
}

const (
	MAX_PLAYERS = 4
)

func NewFarasPlayerManager() *FarasPlayerManager {
	return &FarasPlayerManager{
		Players: make([]*Juwadey, 0),
	}
}

func (pm *FarasPlayerManager) AddPlayer(juwadey *Juwadey) {
	pm.Players = append(pm.Players, juwadey)
}

func (pm *FarasPlayerManager) GetPlayers() []*Juwadey {
	return pm.Players
}

func (pm *FarasPlayerManager) DealCardToPlayer(card *bung.Taas, playerIdx int) {
	pm.Players[playerIdx].Haat = append(pm.Players[playerIdx].Haat, card)
}

func (pm *FarasPlayerManager) CheckIfFull() bool {
	return len(pm.Players) >= MAX_PLAYERS
}

func RotatePlayers(players []*Juwadey, currentPlayerIdx int) []*Juwadey {
	n := len(players)
	rotated := make([]*Juwadey, n)
	for i := range n {
		// Calculate the position based on the current player's perspective
		rotated[i] = players[(currentPlayerIdx+i)%n]
	}
	return rotated
}
