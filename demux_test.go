package containers

import (
	"testing"

	"github.com/julio77it/go-containers/applications/demux"
	"github.com/julio77it/go-containers/set"
	"github.com/stretchr/testify/assert"
)

func TestDemux(t *testing.T) {
	max := 10
	d := make([]set.Set[int], max, max)

	for i := 0; i < max; i++ {
		d[i] = set.New[int]()
	}

	demux := demux.New(
		func(t int, l int) int { return t },
		func(t int, d set.Set[int]) { d.Put(t) },
		d,
	)

	assert.False(t, d[4].Contains(4))
	demux.Put(4)
	assert.True(t, d[4].Contains(4))

	assert.False(t, d[2].Contains(2))
	demux.Put(2)
	assert.True(t, d[2].Contains(2))

	assert.Panics(t, func() { demux.Put(20) })
}
