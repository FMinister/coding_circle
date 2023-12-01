package higherorderfunctions_test

import (
	"testing"

	hof "github.com/FMinister/coding_circle/higherorderfunctions"
	"github.com/stretchr/testify/assert"
)

func TestCdr(t *testing.T) {
	tests := []struct {
		description string
		cons        func(hof.PairFunc) int
		expected    int
	}{
		{
			description: "Cdr of 3, 4",
			cons:        hof.Cons(3, 4),
			expected:    4,
		},
		{
			description: "Cdr of 10, 22",
			cons:        hof.Cons(10, 22),
			expected:    22,
		},
		{
			description: "Cdr of 100, -2",
			cons:        hof.Cons(100, -2),
			expected:    -2,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			assert.Equal(t, test.expected, hof.Cdr(test.cons))
		})
	}
}
