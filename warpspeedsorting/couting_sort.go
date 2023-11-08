package warpspeedsorting

func CountingSort(zipCodes []int) []int {
	return []int{}
}

func FindMax(list []int) int {
	if len(list) == 0 {
		return 0
	}

	max := list[0]

	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}
