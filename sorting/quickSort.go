package sorting

// low --> starting index, high --> ending index

// QuickSort array and low and high index to sort the array...
func QuickSort(arr []int, low int, high int) {

	if low < high {
		// pi is partitioning index , arr[pi] is now at right place
		pi := Partitioning(arr, low, high)

		QuickSort(arr, low, pi-1)  //Before pi
		QuickSort(arr, pi+1, high) //After pi
	}
}

// Partitioning takes array and low and high index and its return index...
func Partitioning(arr []int, low int, high int) int {

	// Pivot element to be placed at right position
	pivot := arr[high]

	i := (low - 1)

	for j := low; j < (high - 1); j++ {

		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return (i + 1)

}
