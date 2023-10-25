package tictactoe_test

import (
	"testing"

	ttt "github.com/FMinister/coding_circle/tictactoe"
	"github.com/stretchr/testify/assert"
)

func TestHasWinningPosition(t *testing.T) {
	tests := []struct {
		description string
		n           int
		input       []ttt.Coordinate
		expected    bool
	}{
		{"single item in list", 3, []ttt.Coordinate{{0, 0}}, false},
		{"three items in list, but no winning position", 3, []ttt.Coordinate{{0, 1}, {1, 1}, {2, 2}}, false},
		{"three items in list, diagonal, left to right, winning position", 3, []ttt.Coordinate{{0, 0}, {1, 1}, {2, 2}}, true},
		{"three items in list, diagonal, right to left, winning position", 3, []ttt.Coordinate{{0, 2}, {1, 1}, {2, 0}}, true},
		{"three items in list, first row winning position", 3, []ttt.Coordinate{{0, 0}, {0, 1}, {0, 2}}, true},
		{"three items in list, first col winning position", 3, []ttt.Coordinate{{0, 0}, {1, 0}, {2, 0}}, true},
		{"four items in list, but no winning position", 3, []ttt.Coordinate{{0, 0}, {1, 1}, {2, 0}, {2, 1}}, false},
		{"five items in list, but no winning position", 3, []ttt.Coordinate{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 1}}, false},
		{"five items in list, row as winning position", 3, []ttt.Coordinate{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}, true},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual := ttt.HasWinningPosition(test.n, test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
