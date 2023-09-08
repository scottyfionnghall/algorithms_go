package main

import (
	"fmt"
	"math/rand"
)

func generateSeq(total int, min int, max int) []int {
	var seq []int
	for i := 0; i <= total; i++ {
		n := rand.Intn(max-min) + min
		seq = append(seq, n)
	}
	return seq
}

func insertionSort(arr []int, n int) {
	for j := 1; j < n; j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = key
	}
	fmt.Println(arr)
}

func main() {
	seq := generateSeq(10000, 0, 10000)
	insertionSort(seq, len(seq))
}
