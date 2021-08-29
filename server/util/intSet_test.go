package util

import (
	"testing"

	"gotest.tools/assert"
)

func TestIntSet(t *testing.T) {
	set := IntSet{}
	set.Add(2, 3, 4)
	t.Run("contains 2", func(t *testing.T) {
		assert.Assert(t, set.Contains(2))
	})
	t.Run("contains 2 3 AND 4", func(t *testing.T) {
		assert.Assert(t, set.ContainsAll(2, 3, 4))
	})
	t.Run("does not contain 0 2 AND 3", func(t *testing.T) {
		assert.Assert(t, set.ContainsAll(0, 2, 3))
	})
	t.Run("contains 1 2 OR 6", func(t *testing.T) {
		assert.Assert(t, set.ContainsAny(1, 2, 6))
	})
	t.Run("does not contain 0 1 OR 5", func(t *testing.T) {
		assert.Assert(t, !set.ContainsAny(0, 1, 5))
	})
	t.Run("2 and 4 correctly deleted", func(t *testing.T) {
		set.Delete(2, 4)
		assert.Assert(t, !set.Contains(2) && !set.Contains(4))
	})
	t.Run("3 was NOT deleted", func(t *testing.T) {
		assert.Assert(t, set.Contains(3))
	})
	t.Run("adding an element twice only increases size by 1", func(t *testing.T) {
		currentSize := len(set)
		set.Add(1)
		set.Add(1)
		assert.Assert(t, len(set) == 1+currentSize)
	})
}
