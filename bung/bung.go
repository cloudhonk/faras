package bung

import "math/rand/v2"

type Bung []Taas

func New() Bung {
	bung := Bung{}
	for _, patti := range Pattiharu {
		for _, rangi := range Rangiharu {
			bung = append(bung, Taas{Patti: patti, Rangi: rangi})
		}
	}
	bung.Fitt()
	return bung
}

func (b *Bung) Fitt() {
	rand.Shuffle(len(*b), func(i, j int) { (*b)[i], (*b)[j] = (*b)[j], (*b)[i] })
}
