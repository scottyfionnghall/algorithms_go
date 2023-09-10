package main

import (
	sort "customsort"
	"fmt"
	"math/rand"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func seqGen(low int, high int) []int {
	var seq []int
	for i := low; i <= high; i++ {
		seq = append(seq, i)
	}
	rand.Shuffle(len(seq), func(i, j int) {
		seq[i], seq[j] = seq[j], seq[i]
	})
	return seq
}

func main() {
	seq := seqGen(0, 1000000)
	defer timer("QuickSort")()
	sort.QuickSort(seq, 0, len(seq)-1)
}
