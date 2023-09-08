package sorting

// SelectionSort sorts an array using selection sort algorithm
// Time complexity: O(n^2)
// Auxiliary space: O(1)
func SelectionSort(arr *[]int, n int) {
	// One by one move boundary of unsorted subarray
	for i := 0; i < n; i++ {
		// Find the minimum element in unsorted array
		minIdx := i
		for j := i + 1; j < n; j++ {
			if (*arr)[j] < (*arr)[minIdx] {
				minIdx = j
			}
		}

		// Swap the found minimum element
		// With the first element
		if minIdx != i {
			(*arr)[minIdx], (*arr)[i] = (*arr)[i], (*arr)[minIdx]
		}
	}
}

// InsertionSort sorts an array using insertion sort algorithm
// Time complexity: O(n^2)
// Auxiliary space: O(1)
func InsertionSort(arr *[]int, n int) {
	for i := 1; i < n; i++ {
		key := (*arr)[i]
		j := i - 1

		// Move elements of arr[0...i-1] that are greater than key,
		// To the one position ahead of their current position
		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

// MergeSort sorts an array using the merge sort algorithm
func MergeSort(arr *[]int) {
	mergeSort(arr, 0, len(*arr)-1)
}

func mergeSort(arr *[]int, l int, r int) {
	if l < r {
		// Find the middle
		m := (l + r) / 2
		// Sort first and second halves
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		// Merge the sorted halves
		merge(arr, l, m, r)
	}
}

// First subarray is arr[1..m]
// Second subarray is arr[m+1...r]
func merge(arr *[]int, l int, m int, r int) {
	// Find sizes of two subarrays to be merged
	n1 := m - l + 1
	n2 := r - m

	// Create temp arrays
	left := make([]int, n1)
	right := make([]int, n2)

	// Copy data to temp arrays
	for i := 0; i < n1; i++ {
		left[i] = (*arr)[l+i]
	}
	for i := 0; i < n2; i++ {
		right[i] = (*arr)[m+1+i]
	}

	// Merge the temp arrays
	i, j, k := 0, 0, l

	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			(*arr)[k] = left[i]
			i++
		} else {
			(*arr)[k] = right[j]
			j++
		}
		k++
	}

	// Copy Remaining elements of left[] if any
	for i < n1 {
		(*arr)[k] = left[i]
		i++
		k++
	}

	// Copy Remaining elements of right[] if any
	for j < n2 {
		(*arr)[k] = right[j]
		j++
		k++
	}
}

// QuickSort sorts an array using the quick sort algorithm
func QuickSort(arr *[]int) {
	quickSort(arr, 0, len(*arr)-1)
}

func quickSort(arr *[]int, low int, high int) {
	if low >= high {
		return
	}

	// pi is paritioning index, arr[p]
	// is now at the right place
	pi := partition(arr, low, high)

	// Separately sort elements before partition and after partition
	quickSort(arr, low, pi-1)
	quickSort(arr, pi+1, high)
}

func partition(arr *[]int, low int, high int) int {
	// Choose the pivot
	pivot := (*arr)[high]

	// Index of smaller element and indicates the right
	// Position of pivot found so far
	i := low - 1

	for j := low; j <= high-1; j++ {
		// If current element is smaller than the pivot
		if (*arr)[j] < pivot {
			// Increment index of smaller element
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	(*arr)[i+1], (*arr)[high] = (*arr)[high], (*arr)[i+1]

	return i + 1
}
