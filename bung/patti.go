package bung

import "fmt"

type Patti int

const (
	Dua = iota + 2
	Tirki
	Chauka
	Panja
	Chhaka
	Satta
	Attha
	Nahar
	Dahar
	Gulam
	Missi
	Bassa
	Ekka
)

var Pattiharu = []Patti{Dua, Tirki, Chauka, Panja, Chhaka, Satta, Attha, Nahar, Dahar, Gulam, Missi, Bassa, Ekka}

func (p *Patti) String() string {
	switch *p {

	case Dahar:
		return "X"

	case Gulam:
		return "J"

	case Missi:
		return "Q"

	case Bassa:
		return "K"

	case Ekka:
		return "A"
	}
	return fmt.Sprintf("%d", *p)
}
