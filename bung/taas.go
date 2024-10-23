package bung

import "fmt"

type Taas struct {
	Patti
	Rangi
}

func (t *Taas) String() string {
	return fmt.Sprintf("%s%s", t.Rangi.String(), t.Patti.String())
}
