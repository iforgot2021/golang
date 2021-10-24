package algorism

func BubbleSort(input []int64) []int64 {
	for i := range input {
		sorted := true
		for j := 0; j+1 < len(input)-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
	return input
}
