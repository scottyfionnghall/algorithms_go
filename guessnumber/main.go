package main

import (
	"fmt"
	"math/rand"
)

// This algorithm reaches a solution afer at most
// k = 1 + log(n) turns, thus for a number between 0 and 1mil
// the solution can be reached in 1 + log(1000000) = 1 + ~20 = ~21
// this is a Logarithmic algorithm, with a notation of O(log n)
func guess(n int, low int, high int) int {
	turns := 0
	for high >= low {
		turns++
		mid := (low + high) / 2
		if mid == n {
			return turns
		} else if mid < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return turns
}

func main() {
	number := rand.Intn(1000000)
	fmt.Println(number)
	fmt.Println(guess(number, 0, 1000000))
}
