# go-containers
## Golang missing containers - a simple implementation

### Set

Set is built over a map[T]*T, with values always nil

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