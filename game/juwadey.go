package game

import (
	"net"

	"github.com/cloudhonk/faras/bung"
)

const (
	BOTTOM = iota
	RIGHT
	TOP
	LEFT
)

type Juwadey struct {
	conn net.Conn
	Name string
	Haat
}

func NewJuwadey(name string, conn net.Conn) *Juwadey {
	j := Juwadey{
		Name: name,
		conn: conn,
		Haat: make([]*bung.Taas, 0),
	}
	return &j
}
