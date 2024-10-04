package game

import (
	"fmt"
	"sync"
	"time"

	"github.com/cloudhonk/faras/bung"
)

type Faras struct {
	id       uint64
	juwadeys []*Juwadey
	update   chan uint64
	end      chan uint64
	mu       *sync.Mutex
}

func newFaras(id uint64, update, end chan uint64) *Faras {
	f := Faras{
		id:       id,
		juwadeys: []*Juwadey{},
		update:   update,
		end:      end,
		mu:       &sync.Mutex{},
	}
	return &f
}

func (f *Faras) addJuwadey(juwadey *Juwadey) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if len(f.juwadeys) < JUWADEYS_PER_GAME {
		f.juwadeys = append(f.juwadeys, juwadey)
		f.update <- f.id

		if len(f.juwadeys) == JUWADEYS_PER_GAME {
			go f.gameLoop()
		}
	}
}

func (f *Faras) gameLoop() {
	deck := bung.New()

	ticker := time.NewTicker(2 * time.Second)

	for i := range CARDS_PER_JUWADEY {
		for j, juwadey := range f.juwadeys {
			juwadey.Haat = append(juwadey.Haat, &deck[j+i*CARDS_PER_JUWADEY])
			f.update <- f.id
			<-ticker.C
		}
	}
	ticker.Stop()
	var juwadeys []Juwadey

	for _, juwadey := range f.juwadeys {
		juwadeys = append(juwadeys, *juwadey)
	}
	winner := determineWinner(juwadeys)
	winnerHandRank := getHandRank(winner.Haat)

	msg := fmt.Sprintf("\n%s wins game with a %s", winner.Name, handRankToStr(winnerHandRank))
	for _, juwadey := range f.juwadeys {
		juwadey.conn.Write([]byte(msg))
	}

	f.end <- f.id
}
