package bung

type Rangi int

const (
	Paan = iota + 1
	Itta
	Chidi
	Hukum
)

var Rangiharu = []Rangi{Paan, Itta, Chidi, Hukum}

func (s *Rangi) String() string {

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
