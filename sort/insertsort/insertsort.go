package insertsort

func InsertSort(arr []int, n int) []int {
	for j := 1; j < n; j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i -= 1
		}
		arr[i+1] = key
	}
	return arr
}
