package prb1to100

import "fmt"

// Mul3or5 - If we list all the natural numbers below 10 that are multiples of 3 or 5,
//			 we get 3, 5, 6 and 9. The sum of these multiples is 23.
//			 Find the sum of all the multiples of 3 or 5 below 1000.
func Mul3or5(max int) string {
	var sum int

	for i := 1; i < max; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return fmt.Sprintf("sum of all the multiples of 3 or 5 below %d is %d", max, sum)
}
