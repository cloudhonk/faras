package khel

import (
	"fmt"
	"os"
	"time"

	"github.com/cloudhonk/faras/bung"
)

//  TODO: Right now its just some random game simulated as faras. Need to do the actual game logic here.

const MAX_ROUND = 10

type Faras struct {
	Juwadeys     []*Juwadey
	CurrentRound int
}

func NewFaras() *Faras {
	f := Faras{
		Juwadeys: []*Juwadey{},
	}
	return &f
}

func (f *Faras) GameLoop() {
	for f.CurrentRound = 1; f.CurrentRound <= MAX_ROUND; f.CurrentRound++ {
		f.dealCards()
		f.findWinner()
		f.updateScoreboard()
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Game has ended. Server is shutting down.")
	os.Exit(0)
}

func (f *Faras) dealCards() {

	deck := bung.New()

	for i, juwadey := range f.Juwadeys {
		juwadey.Taas = deck[i]
	}
}

func (f *Faras) findWinner() {
	var winner *Juwadey
	highestCard := -1

	for _, juwadey := range f.Juwadeys {
		if int(juwadey.Taas.Patti) > highestCard {
			highestCard = int(juwadey.Taas.Patti)
			winner = juwadey
		}
	}

	if winner != nil {
		winner.Wins++
		for _, Juwadey := range f.Juwadeys {
			fmt.Fprintf(Juwadey.Conn, "Round %d Winner: %s with card %s\n", f.CurrentRound, winner.Name, winner.Taas.String())
		}
	}
}

func (f *Faras) updateScoreboard() {
	for _, juwadey := range f.Juwadeys {
		for _, j := range f.Juwadeys {
			fmt.Fprintf(j.Conn, "Scoreboard: %s has %d wins\n", juwadey.Name, juwadey.Wins)
		}
	}
}
