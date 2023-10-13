package median

import (
	"sort"
)

func SingleMedian(input []int) float64 {
	lengthOddOrEven := len(input) % 2

	sort.Ints(input)

	switch lengthOddOrEven {
	case 0:
		return float64(input[len(input)/2-1]+input[len(input)/2]) / 2
	case 1:
		return float64(input[int(len(input)/2)])
	default:
		return 0.0
	}
}

func MultiMedian(input []int) []float64 {
	var result []float64

	for i := range input {
		result = append(result, SingleMedian(input[:i+1]))
	}

	return result
}
