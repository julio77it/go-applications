package demux

import "fmt"

type Demux[T any, D any] struct {
	hash  func(T, int) int
	apply func(T, D)
	demux []D
}

func New[T any, D any](hash func(T, int) int, apply func(T, D), d []D) *Demux[T, D] {
	return &Demux[T, D]{
		hash:  hash,
		apply: apply,
		demux: d,
	}
}

func (d *Demux[T, D]) Put(value T) {
	if idx := d.hash(value, len(d.demux)); (idx < 0) || (idx >= len(d.demux)) {
		panic(fmt.Sprintf("Demux outbound slice limit: %d [0,%d]", idx, len(d.demux)))
	} else {
		d.apply(value, d.demux[idx])
	}
}
