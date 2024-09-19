package bung

import "fmt"

type Taas struct {
	Patti
	Turup
}

func (t *Taas) String() string {
	return fmt.Sprintf("%s को %s", t.Turup.String(), t.Patti.String())
}
