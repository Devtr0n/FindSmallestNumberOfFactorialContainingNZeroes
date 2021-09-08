package main

import "fmt"

// calculate the smallest number of which the factorial contains 'n' zeros
func findSmallestFactorialForNZeros(n int) int {

	// when 'n' equals 1, the easy answer is 5!
	if n == 1 {
		return 5
	}

	// initialize bottom and top for binary search
	low := 0
	high := 5 * n

	// recursive search
	for low < high {

		// binary bit shift
		middle := (low + high) >> 1

		// determine trailing zeroes
		if HasTrailingZeros(middle, n) {
			high = middle
		} else {
			low = middle + 1
		}

	}

	return low
}

// determine number's factorial contains at-least 'n' trailing zeros
func HasTrailingZeros(p int, n int) bool {

	count := 0
	f := 5

	for f <= p {
		count += p / f
		f *= 5
	}

	return (count >= n)

}

func main() {
	fmt.Println(findSmallestFactorialForNZeros(1))    // expected output 5
	fmt.Println(findSmallestFactorialForNZeros(6))    // expected output 25
	fmt.Println(findSmallestFactorialForNZeros(31))   // expected output 125
	fmt.Println(findSmallestFactorialForNZeros(156))  // expected output 625
	fmt.Println(findSmallestFactorialForNZeros(781))  // expected output 3125
	fmt.Println(findSmallestFactorialForNZeros(3906)) // expected output 15625
}
