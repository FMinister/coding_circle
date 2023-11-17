package floydscirclefinding_test

import (
	"testing"

	fcf "github.com/FMinister/coding_circle/floydscirclefinding"
	"github.com/stretchr/testify/assert"
)

func TestFindCircle(t *testing.T) {
	tests := []struct {
		description string
		list        *fcf.Node
		expected    bool
	}{
		{
			description: "No circle",
			list:        fcf.NewNode(1, fcf.NewNode(3, fcf.NewNode(5, fcf.NewNode(7, fcf.NewNode(9, nil))))),
			expected:    false,
		},
		{
			description: "Circle",
			list: func() *fcf.Node {
				node := fcf.NewNode(12, fcf.NewNode(8, fcf.NewNode(17, fcf.NewNode(2, fcf.NewNode(27, fcf.NewNode(9, nil))))))
				node.Next.Next.Next.Next.Next.Next = node.Next.Next.Next
				return node
			}(),
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			assert.Equal(t, test.expected, fcf.FindCircle(test.list))
		})
	}
}
