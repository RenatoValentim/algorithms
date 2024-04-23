package recursive

func BinarySearch(arr []int, val int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	highestIndex := len(arr) - 1
	middleIndex := int(highestIndex / 2)
	guess := arr[middleIndex]
	if guess == val {
		return middleIndex
	}
	if guess > val {
		return BinarySearch(arr[:middleIndex], val)
	}
	if guess < val {
		index := BinarySearch(arr[middleIndex+1:], val)
		if index == -1 {
			return index
		}
		return middleIndex + 1 + index
	}
	return -1
}
