package median

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleMedian(t *testing.T) {
	tests := []struct {
		description string
		input       []int
		expected    float64
	}{
		{"single item in list", []int{1}, 1},
		{"two items in list", []int{1, 2}, 1.5},
		{"three items in list", []int{1, 2, 3}, 2},
		{"nine items in list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
		{"ten items in list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5},
		{"example list", []int{17, 2, 8, 27, 12, 9}, 10.5},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			assert.Equal(t, test.expected, SingleMedian(test.input))
		})
	}
}
