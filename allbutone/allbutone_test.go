package allbutone_test

import (
	"testing"

	abt "github.com/FMinister/coding_circle/allbutone"
	"github.com/stretchr/testify/assert"
)

func TestAllButOne(t *testing.T) {
	var tests = []struct {
		description string
		input       []int
		expected    []int
	}{
		{
			description: "empty input",
			input:       []int{},
			expected:    []int{},
		},
		{
			description: "one element",
			input:       []int{11},
			expected:    []int{1},
		},
		{
			description: "many elements",
			input:       []int{27, 9, 12, 8, 17, 2},
			expected:    []int{29376, 88128, 66096, 99144, 46656, 396576},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			assert.Equal(t, test.expected, abt.AllButOne(test.input))
		})
	}
}
