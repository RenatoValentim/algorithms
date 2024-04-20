package recursive_test

import (
	"testing"

	"github.com/RenatoValentim/algorithms/internal/recursive"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)

	t.Run(`Should return a negative index if array is nil`, func(t *testing.T) {
		output := recursive.BinarySearch(nil, 3)

		assert.Equal(-1, output)
	})

	t.Run(`Should return a negative index if array is empty`, func(t *testing.T) {
		output := recursive.BinarySearch([]int{}, 3)

		assert.Equal(-1, output)
	})
}
