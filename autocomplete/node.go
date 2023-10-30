package autocomplete

type Node struct {
	Children map[rune]*Node
	IsWord   bool
}

// NewNode creates a new node with an empty map of children. IsWord is set to false by default, but is written here for clarity.
func NewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
		IsWord:   false,
	}
}

// GetWords returns a slice of words that are children of the node. It recursively calls itself on each child node, and appends the prefix to the word if the node is a word.
func (node *Node) GetWords(prefix string) []string {
	words := []string{}

	if node.IsWord {
		words = append(words, prefix)
	}

	for char, childNode := range node.Children {
		// This is a recursive call to GetWords. It will keep calling itself until it reaches a node that has no children.
		// childNode.GetWords(prefix+string(char))... takes a slice of strings and appends it to the words slice.
		words = append(words, childNode.GetWords(prefix+string(char))...)
	}

	return words
}
