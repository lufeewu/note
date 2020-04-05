package main

func SplitNumbers(numbers []int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		for numbers[i] >= 0 && i < j {
			i++
		}
		for numbers[j] < 0 && j > i {
			j--
		}
		if i >= j {
			break
		}
		numbers[i], numbers[j] = numbers[j], numbers[i]
		i++
		j--
	}

	j = i
	for j > 0 && numbers[j] < 0 {
		j--
	}
	i = 0
	for i < j {
		for i <= j && numbers[i] != 0 {
			i++
		}
		for i <= j && numbers[j] == 0 {
			j--
		}
		if i <= j {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
			j--
		}
	}

	return numbers
}
