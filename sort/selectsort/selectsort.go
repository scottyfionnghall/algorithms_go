package selectsort

func swap(arr []int, i int, min_idx int) []int {
	key := arr[min_idx]
	arr[min_idx] = arr[i]
	arr[i] = key
	return arr
}

func SelectSort(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}
		if min_idx != i {
			swap(arr, i, min_idx)
		}
	}
	return arr
}
