package game

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

type frameBuilder interface {
	Build(juwadeys []*Juwadey)
	Flush() []byte
}

type farasGameManager struct {
	frameBuilder
	games     []*Faras
	update    chan uint64
	end       chan uint64
	mu        *sync.Mutex
	gameIdGen uint64
}

func NewFarasGameManager(ffb frameBuilder) *farasGameManager {
	return &farasGameManager{
		frameBuilder: ffb,
		games:        make([]*Faras, 0),
		update:       make(chan uint64),
		end:          make(chan uint64),
		mu:           &sync.Mutex{},
	}
}

func (fgm *farasGameManager) Join(conn net.Conn) {

	var name string
	conn.Write([]byte("Enter your name: "))
	fmt.Fscanln(conn, &name)

	juwadey := newJuwadey(name, conn)

	fgm.mu.Lock()
	defer fgm.mu.Unlock()

	for _, game := range fgm.games {
		game.mu.Lock()
		if len(game.juwadeys) < JUWADEYS_PER_GAME {
			game.mu.Unlock()
			game.addJuwadey(juwadey)
			return
		}
		game.mu.Unlock()
	}

	faras := fgm.newGame()
	faras.addJuwadey(juwadey)
	fgm.games = append(fgm.games, faras)

}

func (fgm *farasGameManager) Update() {

	for id := range fgm.update {
		fgm.broadcast(id)
	}

}

func (fgm *farasGameManager) End() {
	for id := range fgm.end {
		fgm.removeGame(id)
	}
}

func (fgm *farasGameManager) broadcast(id uint64) {

	for _, game := range fgm.games {
		if game.id == id {
			for i, juwadey := range game.juwadeys {
				fgm.frameBuilder.Build(rotatePlayers(game.juwadeys, i))
				juwadey.conn.Write(fgm.frameBuilder.Flush())
			}
			return
		}
	}

}

func (fgm *farasGameManager) removeGame(id uint64) {
	fgm.mu.Lock()
	defer fgm.mu.Unlock()
	for i, game := range fgm.games {
		if game.id == id {
			for _, juwadey := range game.juwadeys {
				juwadey.conn.Close()
			}
			fgm.games = append(fgm.games[:i], fgm.games[i+1:]...)
			return
		}
	}
}

func (fgm *farasGameManager) newGame() *Faras {
	id := atomic.AddUint64(&fgm.gameIdGen, 1)
	return newFaras(id, fgm.update, fgm.end)
}
