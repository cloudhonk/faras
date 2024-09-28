package khel

//  TODO: Right now its just some random game simulated as faras. Need to do the actual game logic here.

type Faras struct {
	Juwadeys     []*Juwadey
	CurrentRound int
	End          chan struct{}
}

func NewFaras() *Faras {
	f := Faras{
		Juwadeys: []*Juwadey{},
		End:      make(chan struct{}),
	}
	return &f
}

// func (f *Faras) GameLoop() {
// 	for f.CurrentRound = 1; f.CurrentRound <= 10; f.CurrentRound++ {
// 		f.dealCards()
// 		f.findWinner()
// 		f.updateScoreboard()
// 		time.Sleep(2 * time.Second)
// 	}

// 	fmt.Println("Game has ended.")
// 	f.End <- struct{}{}
// }

// func (f *Faras) dealCards() {

// 	deck := bung.New()

// 	for i, juwadey := range f.Juwadeys {
// 		juwadey.Haat = []*bung.Taas{&deck[i]}
// 	}
// }

// func (f *Faras) findWinner() {
// 	var winner *Juwadey
// 	highestCard := -1

// 	for _, juwadey := range f.Juwadeys {
// 		if int(juwadey.Haat.Patti) > highestCard {
// 			highestCard = int(juwadey.Haat.Patti)
// 			winner = juwadey
// 		}
// 	}

// 	if winner != nil {
// 		winner.Wins++
// 		for _, Juwadey := range f.Juwadeys {
// 			fmt.Fprintf(Juwadey.Conn, "Round %d Winner: %s with card %s\n", f.CurrentRound, winner.Name, winner.Haat.String())
// 		}
// 	}
// }

// func (f *Faras) updateScoreboard() {
// 	for _, juwadey := range f.Juwadeys {
// 		for _, j := range f.Juwadeys {
// 			fmt.Fprintf(j.Conn, "Scoreboard: %s has %d wins\n", juwadey.Name, juwadey.Wins)
// 		}
// 	}
// }

// func (f *Faras) Reset() {

// 	for _, juwadey := range f.Juwadeys {
// 		juwadey.Conn.Close()
// 	}

// 	f.Juwadeys = []*Juwadey{}
// 	f.CurrentRound = 0
// }
