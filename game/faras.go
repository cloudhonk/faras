package game

import (
	"fmt"
	"sync"
	"time"

	"github.com/cloudhonk/faras/bung"
	"github.com/cloudhonk/faras/logger"
)

type faras struct {
	id       uint64
	juwadeys []*Juwadey
	update   chan uint64
	end      chan uint64
	mu       sync.RWMutex
}

func newFaras(id uint64, update, end chan uint64) *faras {
	f := faras{
		id:       id,
		juwadeys: []*Juwadey{},
		update:   update,
		end:      end,
		mu:       sync.RWMutex{},
	}
	return &f
}

func (f *faras) getJuwadeys() []*Juwadey {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.juwadeys
}

func (f *faras) addJuwadey(juwadey *Juwadey) {

	if f.getTotalJuwadeys() < JUWADEYS_PER_GAME {
		logger.Log.Info(fmt.Sprintf("Adding player %s to game %d", juwadey.Name, f.id))
		f.mu.Lock()
		f.juwadeys = append(f.juwadeys, juwadey)
		f.mu.Unlock()
		f.update <- f.id

		if f.getTotalJuwadeys() == JUWADEYS_PER_GAME {
			logger.Log.Info(fmt.Sprintf("Starting game %d", f.id))
			go f.gameLoop()
		}
	}
}

func (f *faras) getTotalJuwadeys() int {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return len(f.juwadeys)
}

func (f *faras) gameLoop() {
	deck := bung.New()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for i := range CARDS_PER_JUWADEY {
		for j, juwadey := range f.getJuwadeys() {
			taas := &deck[j+i*CARDS_PER_JUWADEY]
			juwadey.Haat = append(juwadey.Haat, taas)
			logger.Log.Info(fmt.Sprintf("Game %d: Dealt card %s to juwadey %s", f.id, taas, juwadey.Name))
			f.update <- f.id
			<-ticker.C
		}
	}
	var juwadeys []Juwadey

	for _, juwadey := range f.juwadeys {
		juwadeys = append(juwadeys, *juwadey)
	}
	winner := determineWinner(juwadeys)
	winnerHandRank := getHandRank(winner.Haat)

	msg := fmt.Sprintf("%s wins game with a %s", winner.Name, handRankToStr(winnerHandRank))
	logger.Log.Info(msg)
	for _, juwadey := range f.juwadeys {
		juwadey.conn.Write([]byte("\n" + msg + "\n"))
	}
	f.end <- f.id
}
