package game

import "sort"

func rotatePlayers(players []*Juwadey, currentPlayeridx int) []*Juwadey {
	n := len(players)
	rotated := make([]*Juwadey, n)
	for i := range n {
		rotated[i] = players[(currentPlayeridx+i)%n]
	}
	return rotated
}

func determineWinner(juwadeys []Juwadey) Juwadey {
	sort.Slice(juwadeys, func(i, j int) bool {
		return compareHands(juwadeys[i].Haat, juwadeys[j].Haat) > 0
	})
	return juwadeys[0]
}

func compareHands(h1, h2 Haat) int {
	rank1 := getHandRank(h1)
	rank2 := getHandRank(h2)

	if rank1 != rank2 {
		return rank1 - rank2
	}

	return int(h1.highCard() - h2.highCard())
}

func handRankToStr(rank int) string {
	switch rank {
	case BADHI:
		return "Badhi"
	case JUTT:
		return "Jutt"
	case COLOR:
		return "Color"
	case RUN:
		return "Run"
	case DABLING_RUN:
		return "Dabling Run"
	case TRIAL:
		return "Trial"
	}
	return ""
}
