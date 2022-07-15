# go-applications
[![Go Report Card](https://goreportcard.com/badge/github.com/julio77it/go-applications)](https://goreportcard.com/report/github.com/julio77it/go-applications)

## Applications

### Filter

A simple goroutine/channel based filter
```go
    type Filter[T any, F comparable] struct {
        // ...
        filter set.Set[F]
    }
```

Creating a filter with a custom criteria : for each T it produces a F value
```go
	filter := filter.New(func(input string) string {
		return input + "KEY"
	})
```

Add a filter value
```go
	filter.Add() <- toFilter
```

If f(T) will produce a F value contained in filters entered by Add(), the T object will be given by Get(), otherwise it'll be ignored
```go
	filter.Put() <- expected

	got = <-filter.Get()
```

### Demux
A simple hash based value demultiplexer
Useful to shred data through containers, channels or struct by a hash function

```go
demux := demux.New[T,D](
		// hash function
		func(t T, l T) int { return int(t) % l },
		// apply the choosen D to the current T
		func(t T, d D) { /* d does something with t */ }, 
		// D populated slice (hash func chooses which D in the slice will apply the T value)
		[]D,
	)
```
