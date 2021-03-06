package collector

import (
	"expvar"
)

type stats struct {
	Started            *expvar.String `name:"started"`
	In                 *expvar.Int    `name:"in"`
	Out                *expvar.Int    `name:"out"`
	IteratorCreated    *expvar.String `name:"iterator-created"`
	IteratorCounter    *expvar.Int    `name:"iterator-counter"`
	IteratorErrLast    *expvar.String `name:"iterator-err-last"`
	IteratorErrCounter *expvar.Int    `name:"iterator-err-counter"`
	IteratorTimeout    *expvar.Int    `name:"iterator-timeout"`
}
