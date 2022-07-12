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
