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

	t.Run(`Should return a negative index if value not in array`, func(t *testing.T) {
		output := recursive.BinarySearch([]int{0, 2}, 3)

		assert.Equal(-1, output)
	})

	t.Run(`Should return the index for the searched value when the guess is equals for the value`, func(t *testing.T) {
		output := recursive.BinarySearch([]int{1, 2, 3, 4, 5}, 3)

		assert.Equal(2, output)
	})

	t.Run(`Should return the index for the searched value when the value is in the array but smaller than the guess`, func(t *testing.T) {
		output := recursive.BinarySearch([]int{1, 2, 3, 4, 5}, 2)

		assert.Equal(1, output)
	})

	t.Run(`Should return the index for the searched value when the value is in the array but is greater than the guess`, func(t *testing.T) {
		output := recursive.BinarySearch([]int{1, 2, 3, 4, 5}, 4)

		assert.Equal(3, output)
	})
}
