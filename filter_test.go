package containers

import (
	"fmt"
	"testing"

	"github.com/julio77it/go-containers/filter"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	convert := func(input string) string {
		return input + "KEY"
	}
	expected := "dummy"
	filtered := convert(expected)

	filter := filter.New(convert)

	filter.Add() <- filtered
	filter.Put() <- expected

	got := <-filter.Get()

	assert.Equal(t, expected, got)

	ignored := "nodummy"

	filter.Put() <- ignored
	filter.Put() <- expected

	got = <-filter.Get()

	assert.Equal(t, expected, got)
}

func ExampleFilter() {
	convert := func(input string) string {
		return input + "KEY"
	}
	expected := "dummy"
	filtered := convert(expected)

	filter := filter.New(convert)

	filter.Add() <- filtered
	filter.Put() <- expected

	if got := <-filter.Get(); expected == got {
		fmt.Printf("%s==%s\n", expected, got)
	} else {
		fmt.Printf("%s!=%s\n", expected, got)
	}

	ignored := "nodummy"

	filter.Put() <- ignored
	filter.Put() <- expected

	if got := <-filter.Get(); expected == got {
		fmt.Printf("%s==%s\n", expected, got)
	} else {
		fmt.Printf("%s!=%s\n", expected, got)
	}
	// Output: dummy==dummy
	// dummy==dummy
}

func BenchmarkFilterNoMatch(b *testing.B) {
	value := "Value"

	matcher := func(input string) string {
		return input + "KEY"
	}

	filter := filter.New(matcher)

	for n := 0; n < b.N; n++ {
		filter.Put() <- value
	}
}

func BenchmarkFilterMatch(b *testing.B) {
	value := "Value"

	matcher := func(input string) string {
		return input + "KEY"
	}

	filter := filter.New(matcher)

	match := matcher(value)
	filter.Add() <- match

	go func(ch <-chan string) {
		for {
			select {
			case <-ch:
				break
			}
		}
	}(filter.Get())

	for n := 0; n < b.N; n++ {
		filter.Put() <- value
	}
}

func BenchmarkFilterBufferedMatch(b *testing.B) {
	value := "Value"

	matcher := func(input string) string {
		return input + "KEY"
	}
	bufferSz := 10
	filter := filter.NewBuffered(uint(bufferSz), uint(bufferSz), matcher)

	match := matcher(value)
	filter.Add() <- match

	for i := 0; i < bufferSz; i++ {
		go func(ch <-chan string) {
			for {
				select {
				case <-ch:
					break
				}
			}
		}(filter.Get())
	}

	for n := 0; n < b.N; n++ {
		filter.Put() <- value
	}
}
