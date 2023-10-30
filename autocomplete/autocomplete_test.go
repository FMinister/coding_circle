package autocomplete_test

import (
	"sort"
	"testing"

	ac "github.com/FMinister/coding_circle/autocomplete"
	"github.com/stretchr/testify/assert"
)

func TestNewTrie(t *testing.T) {
	trie := ac.NewTrie()
	assert.NotNil(t, trie)
	assert.NotNil(t, trie.Root)
}

func TestFindAndAdd(t *testing.T) {
	tests := []struct {
		description string
		words       []string
		prefix      string
		expected    []string
	}{
		{"prefix exists", []string{"go", "golang", "gold", "goroutine"}, "gor", []string{"goroutine"}},
		{"prefix exists with multiple words", []string{"go", "golang", "gold", "goroutine"}, "gol", []string{"golang", "gold"}},
		{"prefix does not exist", []string{"go", "golang", "gold", "goroutine"}, "x", []string{}},
		{"prefix is a word", []string{"go", "golang", "gold", "goroutine"}, "golang", []string{"golang"}},
		{"prefix is empty", []string{"go", "golang", "gold", "goroutine"}, "", []string{"go", "golang", "gold", "goroutine"}},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			words := ac.NewTrie()

			for _, word := range test.words {
				words.Add(word)
			}

			result := words.Find(test.prefix)
			sort.Strings(result)

			assert.Equal(t, test.expected, result)
		})
	}
}
