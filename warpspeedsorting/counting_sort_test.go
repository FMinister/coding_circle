package warpspeedsorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	wss "github.com/FMinister/coding_circle/warpspeedsorting"
)

func TestCountingSort(t *testing.T) {
	tests := []struct {
		description    string
		zipCodes       []int
		sortedZipCodes []int
	}{
		{
			description:    "empty list",
			zipCodes:       []int{},
			sortedZipCodes: []int{},
		},
		{
			description:    "single zip code in list",
			zipCodes:       []int{84076},
			sortedZipCodes: []int{84076},
		},
		{
			description:    "three zip codes in list, already sorted",
			zipCodes:       []int{18069, 84076, 93195},
			sortedZipCodes: []int{18069, 84076, 93195},
		},
		{
			description:    "three zip codes in list, not sorted",
			zipCodes:       []int{84076, 93195, 18311},
			sortedZipCodes: []int{18311, 84076, 93195},
		},
		{
			description:    "ten zip codes in list, already sorted",
			zipCodes:       []int{7819, 12437, 28357, 31629, 35457, 61130, 72365, 78333, 85111, 91233},
			sortedZipCodes: []int{7819, 12437, 28357, 31629, 35457, 61130, 72365, 78333, 85111, 91233},
		},
		{
			description:    "ten zip codes in list, not sorted",
			zipCodes:       []int{78333, 61130, 91233, 31629, 12437, 85111, 72365, 35457, 28357, 7819},
			sortedZipCodes: []int{7819, 12437, 28357, 31629, 35457, 61130, 72365, 78333, 85111, 91233},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := wss.CountingSort(test.zipCodes)
			assert.Equal(t, test.sortedZipCodes, result)
		})
	}
}
