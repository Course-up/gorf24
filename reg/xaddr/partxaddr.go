/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

import (
	"github.com/galaktor/gorf24/reg"
	"github.com/galaktor/gorf24/reg/addr"
)

/* has the same 5byte address as its 'root' except for
   the LSByte which differs. When asked will combine
   the root's first 4 MSBytes and append it's own single
   LSByte to form a 5byte address. */
type Partial struct {
	reg.R

	root *Full
}

// TODO: store slice of 5 bytes; use in ReadFrom and WriteTo

func NewPartial(a addr.A, root *Full, lsb byte) *Partial {
	p := &Partial{reg.New(a, reg.NO_MASK), root}
	p.Set(lsb)
	return p
}

// TODO: this certainly can be done more efficiently...
func (r *Partial) Get() A {
	// use New method to force truncate
	return NewFromI((r.root.Get().ToI() << 8) | uint64(r.R.Get()))
}

func (r *Partial) Set(lsb byte) {
	r.R.Set(lsb)
}

// io.ReaderFrom

// io.WriterTo
