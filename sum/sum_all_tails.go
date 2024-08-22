package sum

func SumAllTails(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i := 0; i < lengthOfNumbers; i++ {
		allNumbers := numbersToSum[i]
		if len(allNumbers) == 0 {
			sums[i] = 0
			continue
		}
		tail := allNumbers[1:]

		sums[i] = Sum(tail)
	}

	return sums
}
