package recursive

import "math"

func BinarySearch(arr []int, val int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	lowestIndex := 0
	highestIndex := len(arr) - 1
	middleIndex := int(math.Round(float64(highestIndex+lowestIndex) / 2))
	guess := arr[middleIndex]
	if guess == val {
		return middleIndex
	}
	if guess > val {
		highestIndex = middleIndex - 1
		return BinarySearch(arr[:highestIndex], val)
	}
	if guess < val {
		lowestIndex = middleIndex + 1
		return BinarySearch(arr[lowestIndex:], val)
	}
	return -1
}
