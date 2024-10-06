package game

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"

	"github.com/cloudhonk/faras/logger"
)

type frameBuilder interface {
	Build(juwadeys []*Juwadey)
	Flush() []byte
}

type farasGameManager struct {
	frameBuilder
	games     map[uint64]*faras
	update    chan uint64
	end       chan uint64
	mu        sync.RWMutex
	gameIdGen uint64
}

func NewFarasGameManager(ffb frameBuilder) *farasGameManager {
	return &farasGameManager{
		frameBuilder: ffb,
		games:        make(map[uint64]*faras),
		update:       make(chan uint64, 100),
		end:          make(chan uint64, 100),
		mu:           sync.RWMutex{},
	}
}

func (fgm *farasGameManager) getGames() map[uint64]*faras {
	fgm.mu.RLock()
	defer fgm.mu.RUnlock()
	return fgm.games
}

func (fgm *farasGameManager) Join(conn net.Conn) {

	var name string
	conn.Write([]byte("Enter your name: "))
	fmt.Fscanln(conn, &name)
	juwadey := newJuwadey(name, conn)

	for _, game := range fgm.getGames() {
		if game.getTotalJuwadeys() < JUWADEYS_PER_GAME {
			game.addJuwadey(juwadey)
			return
		}
	}
	faras := fgm.newGame()
	logger.Log.Info(fmt.Sprintf("Creating new game with id: %d", faras.id))
	fgm.mu.Lock()
	defer fgm.mu.Unlock()
	fgm.games[faras.id] = faras
	faras.addJuwadey(juwadey)

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
	game, ok := fgm.games[id]
	if !ok {
		logger.Log.Error(fmt.Sprintf("Game %d not found", id))
		return
	}
	for i, juwadey := range game.getJuwadeys() {
		fgm.frameBuilder.Build(rotatePlayers(fgm.games[id].getJuwadeys(), i))
		if _, err := juwadey.conn.Write(fgm.frameBuilder.Flush()); err != nil {
			logger.Log.Error(fmt.Sprintf("error writing to player %s: %s", juwadey.Name, err))
		}
	}

}

func (fgm *farasGameManager) removeGame(id uint64) {
	game, ok := fgm.games[id]
	if !ok {
		logger.Log.Error(fmt.Sprintf("Game %d not found", id))
		return
	}
	for _, juwadey := range game.getJuwadeys() {
		if err := juwadey.conn.Close(); err != nil {
			logger.Log.Error(fmt.Sprintf("error closing connection: %s", err))
			continue
		}
	}

	fgm.mu.Lock()
	defer fgm.mu.Unlock()
	delete(fgm.games, id)
	logger.Log.Info(fmt.Sprintf("Game %d ended", id))

}

func (fgm *farasGameManager) newGame() *faras {
	id := atomic.AddUint64(&fgm.gameIdGen, 1)
	return newFaras(id, fgm.update, fgm.end)
}
