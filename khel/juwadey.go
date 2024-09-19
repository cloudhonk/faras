package khel

import (
	"net"

	"github.com/cloudhonk/faras/bung"
)

type Juwadey struct {
	Conn net.Conn
	Name string
	Taas bung.Taas
	Wins int
}

func NewJuwadey(name string, conn net.Conn) *Juwadey {
	j := Juwadey{
		Name: name,
		Conn: conn,
	}
	return &j
}
