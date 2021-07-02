package myarray

func Sum(numbers []int) int {
	sum := 0
	// range return index and value
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numberToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums

}

func SumAllTails(numberToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numberToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)

		} else {
			sums = append(sums, Sum(numbers[1:]))
		}

	}
	return sums
}
