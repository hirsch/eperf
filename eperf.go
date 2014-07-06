//eperf offers an easy solution to analyse the basic performance outline
//of program parts.

package eperf

import (
	"log"
	"time"
)

type Perftest struct {
	name   string
	cycles int
	round  int
	min    time.Duration
	last   time.Time
}

//New creates a new Perftest with a given name and the number of times
//(cycles > 0) the program part will be executed.
func New(name string, cycles int) *Perftest {
	p := &Perftest{name: name, cycles: cycles, round: 0}
	return p
}

//Run runs the given Perftest. Use within a for loop.
//The fastest execution time of all cycles will be printed.
//Example:
//
//perftest := eperf.New("code test", 1000)
//for perftest.Run() {
//		code()
//}
func (p *Perftest) Run() bool {
	if p.round > p.cycles {
		if p.round > 1 {
			log.Println("eperf:", p.name, "runtime:", p.min)
			return false
		} else {
			log.Println("eperf:", p.name, "no cycles!")
			return false
		}
	}
	run := time.Since(p.last)
	if run < p.min || p.round == 0 {
		p.min = run
	}
	p.round++
	p.last = time.Now()
	return true
}
