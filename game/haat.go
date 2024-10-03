package game

import (
	"sort"

	"github.com/cloudhonk/faras/bung"
)

type Haat []*bung.Taas

func (h Haat) isTrail() bool {
	return h[0].Patti == h[1].Patti && h[1].Patti == h[2].Patti
}

func (h Haat) isDablingRun() bool {
	sort.Slice(h, func(i, j int) bool {
		return h[i].Patti < h[j].Patti
	})
	return h[0].Rangi == h[1].Rangi && h[1].Rangi == h[2].Rangi &&
		h[1].Patti == h[0].Patti+1 && h[2].Patti == h[1].Patti+1
}

func (h Haat) isRun() bool {
	sort.Slice(h, func(i, j int) bool {
		return h[i].Patti < h[j].Patti
	})
	return h[1].Patti == h[0].Patti+1 && h[2].Patti == h[1].Patti+1
}

func (h Haat) isColor() bool {
	return h[0].Rangi == h[1].Rangi && h[1].Rangi == h[2].Rangi
}

func (h Haat) isJutt() bool {
	return h[0].Patti == h[1].Patti || h[1].Patti == h[2].Patti || h[0].Patti == h[2].Patti
}

func (h Haat) highCard() bung.Patti {
	highCard := h[0].Patti
	for _, card := range h[1:] {
		if card.Patti > highCard {
			highCard = card.Patti
		}
	}
	return highCard
}

func getHandRank(h Haat) int {
	if h.isTrail() {
		return TRIAL
	}
	if h.isDablingRun() {
		return DABLING_RUN
	}
	if h.isRun() {
		return RUN
	}
	if h.isColor() {
		return COLOR
	}
	if h.isJutt() {
		return JUTT
	}
	return BADHI
}
