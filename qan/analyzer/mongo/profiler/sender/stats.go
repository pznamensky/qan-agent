package sender

import (
	"expvar"
)

type stats struct {
	Started *expvar.String `name:"started"`
	In      *expvar.Int    `name:"in"`
	Out     *expvar.Int    `name:"out"`
	ErrIter *expvar.Int    `name:"err-iter"`
}
