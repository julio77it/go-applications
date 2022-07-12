package set

import (
	"fmt"
)

type Set[T comparable] interface {
	Put(value T)
	Contains(value T) bool
	Remove(value T)
	Size() int
	Entries() []T
}

type setByMap[T comparable] struct {
	container map[T]*T
}

// New init a Set instance
func New[T comparable]() Set[T] {
	return &setByMap[T]{
		container: make(map[T]*T),
	}
}

// Put put value in the container, if needed
func (s *setByMap[T]) Put(value T) {
	if !s.Contains(value) {
		s.container[value] = nil
	}
}

// Contains check value existence in container
func (s *setByMap[T]) Contains(value T) bool {
	_, ok := s.container[value]
	return ok
}

// Remove remove value from the container, if needed
func (s *setByMap[T]) Remove(value T) {
	delete(s.container, value)
}

// Size count entries in the container
func (s *setByMap[T]) Size() int {
	return len(s.container)
}

// Size count entries in the container
func (s *setByMap[T]) Entries() []T {
	keys := make([]T, len(s.container))

	i := 0
	for k := range s.container {
		keys[i] = k
		i++
	}
	return keys
}

// String convert set to string (fmt.Print*)
func (s *setByMap[T]) String() string {
	var result string = "{"
	for k, e := range s.Entries() {
		if k > 0 {
			result += ", "
		}
		result += fmt.Sprint("[", e, "]")
	}
	return result + "}"
}
