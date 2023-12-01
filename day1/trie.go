package day1

type TrieNode struct {
	Children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
	}
}

func (t *TrieNode) Insert(word string) {
	current := t
	for _, c := range word {
		if _, ok := current.Children[c]; !ok {
			current.Children[c] = NewTrieNode()
		}
		current = current.Children[c]
	}
	current.isEnd = true
}

func (t *TrieNode) Search(word string) bool {
	current := t
	for _, c := range word {
		if _, ok := current.Children[c]; !ok {
			return false
		}
		current = current.Children[c]
	}
	return current.isEnd
}