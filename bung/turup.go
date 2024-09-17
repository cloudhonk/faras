package bung

type Turup int

const (
	Paan = iota + 1
	Itta
	Chidi
	Hukum
)

var Turupharu = []Turup{Paan, Itta, Chidi, Hukum}

func (s *Turup) String() string {

	switch *s {

	case Paan:
		return "♥"

	case Itta:
		return "♦"

	case Chidi:
		return "♣"

	case Hukum:
		return "♠"
	}

	return ""

}
