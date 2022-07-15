package filter

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Filter[T any, F comparable] struct {
	add    chan F
	remove chan F
	in     chan T
	out    chan T

	filter  mapset.Set[F]
	convert func(T) F
}

func (this *Filter[T, F]) run() {
	for {
		select {
		case f := <-this.add:
			this.filter.Add(f)

		case f := <-this.remove:
			this.filter.Remove(f)

		case t := <-this.in:
			f := this.convert(t)
			if !this.filter.Contains(f) {
				break
			}
			this.out <- t
		}
	}
}

func (this *Filter[T, F]) Add() chan<- F {
	return this.add
}
func (this *Filter[T, F]) Remove() chan<- F {
	return this.remove
}
func (this *Filter[T, F]) Put() chan<- T {
	return this.in
}
func (this *Filter[T, F]) Get() <-chan T {
	return this.out
}

func New[T any, F comparable](convertTtoF func(T) F) *Filter[T, F] {
	return NewBuffered(0, 0, convertTtoF)
}

func NewBuffered[T any, F comparable](putBuffer uint, getBuffer uint, convertTtoF func(T) F) *Filter[T, F] {
	this := &Filter[T, F]{
		convert: convertTtoF,
		filter:  mapset.NewSet[F](),
		add:     make(chan F),
		remove:  make(chan F),
		in:      make(chan T, putBuffer),
		out:     make(chan T, getBuffer),
	}
	go this.run()

	return this
}
