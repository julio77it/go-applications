# go-containers
[![Go Report Card](https://goreportcard.com/badge/github.com/julio77it/go-containers)](https://goreportcard.com/report/github.com/julio77it/go-containers)

## Golang missing containers - a simple implementation

### Set

Set is built over a map[T]*T, which values will be always nil

Put method checks if item is already in set, otherwise it makes the insert.
This way seems to perfom better than to replace the same item

```
goos: darwin
goarch: amd64
pkg: github.com/julio77it/go-containers
cpu: Intel(R) Core(TM) i5-3210M CPU @ 2.50GHz
BenchmarkSetPut-4                       100000000               10.48 ns/op
BenchmarkSetPutWithoutCheck-4           49628702                22.81 ns/op
PASS
ok      github.com/julio77it/go-containers      3.449s
```

### Filter

A simple goroutine/channel based filter
It's a go-containers/set's application
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

