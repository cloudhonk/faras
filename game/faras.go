package game

import (
	"fmt"
	"time"

	"github.com/cloudhonk/faras/bung"
	"github.com/cloudhonk/faras/logger"
)

type faras struct {
	id           uint64
	juwadeys     []*Juwadey
	frameBuilder frameBuilder
}

func newFaras(id uint64) *faras {
	farasFrameConfig := FarasFrameConfig{
		Width:   WINDOW_WIDTH,
		Height:  WINDOW_HEIGHT,
		Padding: WINDOW_PADDING,
	}
	f := faras{
		id:           id,
		juwadeys:     []*Juwadey{},
		frameBuilder: NewFarasFrameBuilder(farasFrameConfig),
	}
	return &f
}

func (f *faras) getJuwadeys() []*Juwadey {
	return f.juwadeys
}

func (f *faras) addJuwadey(juwadey *Juwadey) {

	logger.Log.Info(fmt.Sprintf("Adding player %s to game %d", juwadey.Name, f.id))
	f.juwadeys = append(f.juwadeys, juwadey)
	for i, juwadey := range f.getJuwadeys() {
		f.frameBuilder.Build(rotatePlayers(f.getJuwadeys(), i))
		if _, err := juwadey.conn.Write(f.frameBuilder.Flush()); err != nil {
			logger.Log.Error(fmt.Sprintf("error writing to player %s: %s", juwadey.Name, err))
		}
	}
}

func (f *faras) GameLoop() {
	deck := bung.New()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for i := range CARDS_PER_JUWADEY {
		for j, juwadey := range f.getJuwadeys() {
			taas := &deck[j+i*CARDS_PER_JUWADEY]
			juwadey.Haat = append(juwadey.Haat, taas)
			logger.Log.Info(fmt.Sprintf("Game %d: Dealt card %s to juwadey %s", f.id, taas, juwadey.Name))
			for i, juwadey := range f.getJuwadeys() {
				f.frameBuilder.Build(rotatePlayers(f.getJuwadeys(), i))
				if _, err := juwadey.conn.Write(f.frameBuilder.Flush()); err != nil {
					logger.Log.Error(fmt.Sprintf("error writing to player %s: %s", juwadey.Name, err))
				}
			}
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
		juwadey.conn.Close()
	}
}
