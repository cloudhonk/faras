package khel

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
	Conn net.Conn
	Name string
	Haat []*bung.Taas
	Wins int
}

func NewJuwadey(name string, conn net.Conn) *Juwadey {
	j := Juwadey{
		Name: name,
		Conn: conn,
	}
	return &j
}
