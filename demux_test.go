package containers

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/julio77it/go-applications/demux"
	"github.com/stretchr/testify/assert"
)

func TestDemux(t *testing.T) {
	max := 10
	d := make([]mapset.Set[int], max, max)

	for i := 0; i < max; i++ {
		d[i] = mapset.NewSet[int]()
	}
	demux := demux.New(
		func(t int, l int) int { return t },
		func(t int, d mapset.Set[int]) { d.Add(t) },
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
