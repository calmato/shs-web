package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Reset(t *testing.T) {
	t.Parallel()
	s := New(10)
	s.Add(1)
	s.Reset(10)
	assert.False(t, s.Contains(1))
}

func TestSet_Len(t *testing.T) {
	t.Parallel()
	s := New(10)
	s.Add(1)
	assert.Equal(t, 1, s.Len())
}

func TestSet_Contains(t *testing.T) {
	t.Parallel()
	s := New(10)
	s.Add(1)
	s.Add("test")
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains("test"))
	assert.False(t, s.Contains(2))
	assert.False(t, s.Contains("debug"))
}

func TestSet_FindOrAdd(t *testing.T) {
	t.Parallel()
	s := New(10)
	assert.True(t, s.FindOrAdd(1))
	assert.True(t, s.FindOrAdd("test"))
	assert.False(t, s.FindOrAdd(1))
	assert.False(t, s.FindOrAdd("test"))
}

func TestSet_Remove(t *testing.T) {
	t.Parallel()
	s := New(10)
	s.Add(1)
	assert.True(t, s.Contains(1))
	s.Remove(1)
	assert.False(t, s.Contains(1))
}

func TestSet_Do(t *testing.T) {
	t.Parallel()
	s := New(10)
	target := []int64{1, 2, 3, 4}
	for _, n := range target {
		s.Add(n)
	}
	s.Do(func(v interface{}) {
		actual, ok := v.(int64)
		assert.True(t, ok)
		assert.Contains(t, target, actual)
	})
}

func TestSet_Strings(t *testing.T) {
	t.Parallel()
	s := New(10)
	s.AddStrings([]string{"a", "c", "b"}...)
	assert.Equal(t, []string{"a", "b", "c"}, s.SortStrings())
}

func TestSet_Int64s(t *testing.T) {
	t.Parallel()
	s := New(10)
	s.AddInt64s([]int64{1, 3, 2}...)
	assert.Equal(t, []int64{1, 2, 3}, s.SortInt64s())
}
