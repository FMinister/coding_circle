package continuous_max

func GetMax(values []int, k int) []int {
	maxValues := make([]int, 0, len(values)-k+1)

	var window []int
	for i := 0; i < len(values); i++ {
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}

		for len(window) > 0 && values[window[len(window)-1]] <= values[i] {
			window = window[:len(window)-1]
		}

		window = append(window, i)

		if i >= k-1 {
			maxValues = append(maxValues, values[window[0]])
		}
	}

	return maxValues
}

func GetMaxOther(values []int, k int) []int {
	maxValues := make([]int, 0, len(values)-k+1)

	var window []int
	for i := 0; i < len(values); i++ {
		window = append(window, values[i])

		if len(window) > k {
			window = window[1:]
		}

		if len(window) == k {
			maxValue := 0
			for _, v := range window {
				if v > maxValue {
					maxValue = v
				}
			}

			maxValues = append(maxValues, maxValue)
		}
	}

	return maxValues
}
