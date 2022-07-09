package containers

import (
	"fmt"
	"testing"

	"github.com/julio77it/go-containers/set"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	container := set.New[string]()

	value1 := "Value1"
	value2 := "Value2"
	value3 := "Value3"

	container.Put(value1)
	assert.Truef(t, container.Contains(value1), "Missing entry : %s", value1)

	container.Put(value2)
	assert.Truef(t, container.Contains(value1), "Missing entry : %s", value1)
	assert.Truef(t, container.Contains(value2), "Missing entry : %s", value2)

	container.Put(value3)
	assert.Truef(t, container.Contains(value1), "Missing entry : %s", value1)
	assert.Truef(t, container.Contains(value2), "Missing entry : %s", value2)
	assert.Truef(t, container.Contains(value3), "Missing entry : %s", value3)

	container.Remove(value2)
	assert.Truef(t, container.Contains(value1), "Missing entry : %s", value1)
	assert.Falsef(t, container.Contains(value2), "Found entry : %s", value2)
	assert.Truef(t, container.Contains(value3), "Missing entry : %s", value3)
}

func BenchmarkSetPut(b *testing.B) {
	value := "Value"

	container := set.New[string]()

	for n := 0; n < b.N; n++ {
		container.Put(value)
	}
}

func ExamplePut1() {
	s := set.New[string]()
	s.Put("Entry1")
	fmt.Println("Entries :", s)
	// Output: Entries : {[Entry1]}
}

func ExampleRemove1() {
	s := set.New[string]()
	s.Put("Entry1")
	s.Put("Entry2")
	s.Remove("Entry2")
	fmt.Println("Entries :", s)
	// Output: Entries : {[Entry1]}
}

func ExamplePutAndRemove100() {
	s := set.New[string]()

	s.Put("EntryX")

	for i := 1; i <= 10; i++ {
		s.Put(fmt.Sprintf("Entry%d", i))
	}
	for i := 1; i <= 10; i++ {
		s.Remove(fmt.Sprintf("Entry%d", i))
	}
	fmt.Println("Entries :", s)
	// Output: Entries : {[EntryX]}
}
