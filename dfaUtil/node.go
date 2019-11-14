package dfaUtil

const (
	INIT_TRIE_CHILDREN_NUM = 128	// Since we need to deal all kinds of language, so we use 128 instead of 26
)

// TrieNode data structure
// TrieNode itself doesn't have any value. The value is represented on the path
type TrieNode struct {
	// if this node is the end of a word
	isEndOfWord bool

	// the collection of children of this node
	children map[rune]*TrieNode
}

// Create new TrieNode
func newTrieNode() *TrieNode {
	return &TrieNode{
		isEndOfWord: false,
		children: make(map[rune]*TrieNode, INIT_TRIE_CHILDREN_NUM),
	}
}
