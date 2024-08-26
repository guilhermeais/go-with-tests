package sum

import "fmt"

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllWithDynamicCap(numbersToSum ...[]int) []int {
	sums := []int{}
	fmt.Printf("len() = %d, cap() = %d\n", len(sums), cap(sums))

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
		fmt.Printf("len() = %d, cap() = %d\n", len(sums), cap(sums))

	}

	return sums
}
