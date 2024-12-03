package game

import (
	"sync/atomic"
)

type frameBuilder interface {
	Build(juwadeys []*Juwadey)
	Flush() []byte
}

type farasGameManager struct {
	gameIdGen uint64
}

func NewFarasGameManager() *farasGameManager {
	return &farasGameManager{}
}

func (fgm *farasGameManager) Begin(juwadeyChan chan *Juwadey) {
	count := 0
	var game *faras
	for juwadey := range juwadeyChan {
		count++
		if game == nil {
			game = fgm.newGame()
		}
		game.addJuwadey(juwadey)

		if count == JUWADEYS_PER_GAME {
			count = 0
			go game.GameLoop()
			game = nil
		}
	}
}

func (fgm *farasGameManager) newGame() *faras {
	id := atomic.AddUint64(&fgm.gameIdGen, 1)
	return newFaras(id)
}
