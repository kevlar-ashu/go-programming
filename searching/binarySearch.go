package searching

// BinarySearch will search the key from item array and return boolean value...
func BinarySearch(key int, item []int) bool {

	low := 0
	high := len(item) - 1

	for low <= high {
		median := (high + low) / 2

		if item[median] == key {
			return true
		} else if item[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	return false
}
