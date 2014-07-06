// Package eperf offers an easy solution to analyse the basic performance outline
// of program parts.
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

// New creates a new Perftest with a given name and the number of times
// (cycles > 0) the program part will be executed.
func New(name string, cycles int) *Perftest {
	return &Perftest{name: name, cycles: cycles, round: 0}
}

// Run runs the given Perftest. Use within a for loop.
// The fastest execution time of all cycles will be printed.
//	perftest := eperf.New("code test", 1000)	// Create a new Perftest
//	for perftest.Run() {				// Use Run() in a loop
//		code()					// code() will be executed 1000 times.
//	}
//	// The fastest execution time will be logged. Example output:
//	// 2014/07/06 17:26:59 eperf: code test runtime: 133.806us
func (p *Perftest) Run() bool {
	if p.round > p.cycles {
		if p.round > 1 {
			log.Println("eperf:", p.name, "runtime:", p.min)
			return false
		}
		log.Println("eperf:", p.name, "no cycles!")
		return false
	}
	run := time.Since(p.last)
	if run < p.min || p.round == 0 {
		p.min = run
	}
	p.round++
	p.last = time.Now()
	return true
}
