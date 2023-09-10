package quicksort

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
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
