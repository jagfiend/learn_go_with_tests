package slice_n_dice

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var totes []int

	for _, numbers := range numbersToSum {
		totes = append(totes, Sum(numbers))
	}

	return totes
}

func SumAllTails(numbersToTail ...[]int) []int {
	var tails []int

	for _, numbers := range numbersToTail {
		if len(numbers) == 0 {
			tails = append(tails, 0)
		} else {
			tail := numbers[1:]
			tails = append(tails, Sum(tail))
		}
	}

	return tails
}
