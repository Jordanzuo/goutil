package dfaUtil

const (
	INIT_TRIE_CHILDREN_NUM = 0
)

// trieNode data structure
// trieNode itself doesn't have any value. The value is represented on the path
type trieNode struct {
	// if this node is the end of a word
	isEndOfWord bool

	// the collection of children of this node
	children map[rune]*trieNode
}

// Create new trieNode
func newtrieNode() *trieNode {
	return &trieNode{
		isEndOfWord: false,
		children:    make(map[rune]*trieNode, INIT_TRIE_CHILDREN_NUM),
	}
}
