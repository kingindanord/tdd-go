package matho

func Sum(numbers []int) int {

	var sum int
	for _, v := range numbers {
		sum += v
	}

	return sum

}

func SumAll(nombersToSum ...[]int) (sums []int) {

	for _, v := range nombersToSum {

		sums = append(sums, Sum(v))

	}

	return sums

}

// beter performance
func SumAll2(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {

			sums[i] = Sum(numbers[1:])
		}
	}
	return sums
}
