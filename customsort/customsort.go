package customsort

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func SelectSort(array []int, n int) {
	for i := 0; i < n-1; i++ {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min_idx] {
				min_idx = j
			}
		}
		if min_idx != i {
			swap(&array[i], &array[min_idx])
		}
	}
}

func InsertSort(array []int, n int) {
	for j := 1; j < n; j++ {
		key := array[j]
		i := j - 1
		for i >= 0 && array[i] > key {
			array[i+1] = array[i]
			i -= 1
		}
		array[i+1] = key
	}
}

func partition(array []int, low int, high int) int {
	pivot := array[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if array[j] < pivot {
			i++
			swap(&array[i], &array[j])
		}
	}
	swap(&array[i+1], &array[high])
	return i + 1
}

func QuickSort(array []int, low int, high int) {
	if low < high {
		pi := partition(array, low, high)
		QuickSort(array, low, pi-1)
		QuickSort(array, pi+1, high)
	}
}
