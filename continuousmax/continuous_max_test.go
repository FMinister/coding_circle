package continuous_max_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cm "github.com/FMinister/coding_circle/continuousmax"
)

func TestGetMax(t *testing.T) {
	tests := []struct {
		description string
		values      []int
		k           int
		expected    []int
	}{
		{"get max with k = 3 and list.length = 6", []int{27, 9, 17, 2, 12, 8}, 3, []int{27, 17, 17, 12}},
		{"get max with k = 2 and list.length = 6", []int{27, 9, 17, 2, 12, 8}, 2, []int{27, 17, 17, 12, 12}},
		{"get max with k = 4 and list.length = 10", []int{27, 9, 17, 2, 12, 8, 1, 22, 20, 44}, 10, []int{27, 17, 17, 12, 22, 22, 44}},
		{"get max with k = 5 and list.length = 10", []int{27, 9, 17, 2, 12, 8, 1, 22, 20, 44}, 10, []int{27, 17, 17, 22, 22, 44}},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual := cm.GetMax(test.values, test.k)
			assert.Equal(t, test.expected, actual)
		})
	}
}
