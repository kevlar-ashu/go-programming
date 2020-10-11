package searching

func LinearSearch(items []int, data int) bool {
	for _, key := range items {
		if key == data {
			return true
		}
	}
	return false
}
