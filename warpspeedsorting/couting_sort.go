package warpspeedsorting

func CountingSort(zipCodes []int) []int {
	if len(zipCodes) == 0 || len(zipCodes) == 1 {
		return zipCodes
	}

	max, min := FindMaxAndMin(zipCodes)

	countingArray := make([]int, max-min+1)

	for _, zipCode := range zipCodes {
		countingArray[zipCode-min]++
	}

	sortedZipCodes := make([]int, 0, len(zipCodes))

	for i, count := range countingArray {
		for count > 0 {
			sortedZipCodes = append(sortedZipCodes, i+min)
			count--
		}
	}

	return sortedZipCodes
}

func FindMaxAndMin(list []int) (int, int) {
	if len(list) == 0 {
		return 0, 0
	}

	max, min := list[0], list[0]

	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
		if list[i] < min {
			min = list[i]
		}
	}

	return max, min
}
