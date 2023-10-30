package autocomplete

// Root:
//
//	Children:
//	  c:
//	    Children:
//	      a:
//	        Children:
//	          t:
//	            Children: {}
//	            IsWord: true
type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
	}
}

// Add adds a word to the trie. It iterates through each character in the word, and adds a new node to the trie if it doesn't already exist.
func (trie *Trie) Add(word string) {
	currentNode := trie.Root

	for _, char := range word {
		if _, ok := currentNode.Children[char]; !ok {
			currentNode.Children[char] = NewNode()
		}

		currentNode = currentNode.Children[char]
	}

	currentNode.IsWord = true
}

// Same as Add but without adding a new node if it doesn't exist. ;)
func (trie *Trie) Find(prefix string) []string {
	currentNode := trie.Root

	for _, char := range prefix {
		if _, ok := currentNode.Children[char]; !ok {
			return []string{}
		}

		currentNode = currentNode.Children[char]
	}

	return currentNode.GetWords(prefix)
}
