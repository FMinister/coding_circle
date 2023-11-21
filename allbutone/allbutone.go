package allbutone

func AllButOne(input []int) []int {
	if len(input) < 1 {
		return input
	}
	if len(input) == 1 {
		return []int{1}
	}

	output := make([]int, len(input))
	for i := range input {
		output[i] = 1
	}

	leftProduct := 1
	for i := 0; i < len(input); i++ {
		output[i] = output[i] * leftProduct
		leftProduct = leftProduct * input[i]
	}

	rightProduct := 1
	for i := len(input) - 1; i >= 0; i-- {
		output[i] = output[i] * rightProduct
		rightProduct = rightProduct * input[i]
	}

	return output
}
