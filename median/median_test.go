package median_test

import (
	"testing"

	"github.com/FMinister/coding_circle/median"
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
			result := median.SingleMedian(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestMultiMedian(t *testing.T) {
	tests := []struct {
		description string
		input       []int
		expected    []float64
	}{
		{"single item in list", []int{1}, []float64{1}},
		{"two items in list", []int{1, 2}, []float64{1, 1.5}},
		{"three items in list", []int{1, 2, 3}, []float64{1, 1.5, 2}},
		{"nine items in list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}},
		{"ten items in list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5}},
		{"example list", []int{17, 2, 8, 27, 12, 9}, []float64{17, 9.5, 8, 12.5, 12, 10.5}},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := median.MultiMedian(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestRecursiveMedian(t *testing.T) {
	tests := []struct {
		description string
		input       []int
		expected    []float64
	}{
		{"single item in list", []int{1}, []float64{1}},
		{"two items in list", []int{1, 2}, []float64{1, 1.5}},
		{"three items in list", []int{1, 2, 3}, []float64{1, 1.5, 2}},
		{"nine items in list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}},
		{"ten items in list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5}},
		{"example list", []int{17, 2, 8, 27, 12, 9}, []float64{17, 9.5, 8, 12.5, 12, 10.5}},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := median.RecursiveMedian(test.input, []float64{}, 0)
			assert.Equal(t, test.expected, result)
		})
	}
}
