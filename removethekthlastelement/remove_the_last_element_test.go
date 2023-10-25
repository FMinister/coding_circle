package removethelastelement_test

import (
	"testing"

	rtle "github.com/FMinister/coding_circle/removethekthlastelement"
	"github.com/stretchr/testify/assert"
)

func TestRemoveTheLastElement(t *testing.T) {
	t.Run("Remove the last element", func(t *testing.T) {
		list := rtle.NewNode(1, rtle.NewNode(3, rtle.NewNode(5, rtle.NewNode(7, rtle.NewNode(9, nil)))))
		updatedList := rtle.RemoveTheLastElement(list, 1)
		values := rtle.GetListValues(updatedList)

		assert.Equal(t, []int{1, 3, 5, 7}, values)
	})

	t.Run("Remove the first element", func(t *testing.T) {
		list := rtle.NewNode(1, rtle.NewNode(3, rtle.NewNode(5, rtle.NewNode(7, rtle.NewNode(9, nil)))))
		updatedList := rtle.RemoveTheLastElement(list, 5)
		values := rtle.GetListValues(updatedList)

		assert.Equal(t, []int{3, 5, 7, 9}, values)
	})

	t.Run("Remove the i-th element", func(t *testing.T) {
		list := rtle.NewNode(1, rtle.NewNode(3, rtle.NewNode(5, rtle.NewNode(7, rtle.NewNode(9, nil)))))
		updatedList := rtle.RemoveTheLastElement(list, 3)
		values := rtle.GetListValues(updatedList)

		assert.Equal(t, []int{1, 3, 7, 9}, values)
	})

	t.Run("Remove nothing because i is too large", func(t *testing.T) {
		list := rtle.NewNode(1, rtle.NewNode(3, rtle.NewNode(5, rtle.NewNode(7, rtle.NewNode(9, nil)))))
		updatedList := rtle.RemoveTheLastElement(list, 6)

		assert.Nil(t, updatedList)
	})
}
