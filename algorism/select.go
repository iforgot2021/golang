package algorism

func SelectionSort(input []int64) []int64 {
	for i := range input {
		idx, mVal := i, input[i]
		for j := i + 1; j < len(input); j++ {
			if mVal > input[j] {
				idx, mVal = j, input[j]
			}
		}

		if idx != i {
			input[i], input[idx] = input[idx], input[i]
		}
	}
	return input
}
