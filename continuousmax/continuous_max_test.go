package continuous_max_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	cm "github.com/FMinister/coding_circle/continuousmax"
)

var tests = []struct {
	description string
	values      []int
	k           int
	expected    []int
}{
	{"get max with k = 3 and list.length = 6", []int{27, 9, 17, 2, 12, 8}, 3, []int{27, 17, 17, 12}},
	{"get max with k = 2 and list.length = 6", []int{27, 9, 17, 2, 12, 8}, 2, []int{27, 17, 17, 12, 12}},
	{"get max with k = 4 and list.length = 10", []int{27, 9, 17, 2, 12, 8, 1, 22, 20, 44}, 4, []int{27, 17, 17, 12, 22, 22, 44}},
	{"get max with k = 5 and list.length = 10", []int{27, 9, 17, 2, 12, 8, 1, 22, 20, 44}, 5, []int{27, 17, 17, 22, 22, 44}},
}

func TestGetMax(t *testing.T) {
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual := cm.GetMax(test.values, test.k)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestGetMaxOther(t *testing.T) {
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual := cm.GetMaxOther(test.values, test.k)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func BenchmarkGetMax(b *testing.B) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("Benchmarking: %s", test.description), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cm.GetMax(test.values, test.k)
			}
		})
	}
}

func BenchmarkGetMaxOther(b *testing.B) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("Benchmarking: %s", test.description), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cm.GetMaxOther(test.values, test.k)
			}
		})
	}
}

// go test ./continuousmax -bench=. -benchmem
