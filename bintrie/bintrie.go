package main

import (
	"fmt"
)

type BinTrie struct {
	Value string
	Left  *BinTrie
	Right *BinTrie
}

// Dfs does the inorder traversal for a binary trie
func (bt *BinTrie) Dfs(f func(*BinTrie)) {
	if bt == nil {
		return
	}
	bt.Left.Dfs(f)
	f(bt)
	bt.Right.Dfs(f)

}

func createTestTrie() *BinTrie {
	var root = &BinTrie{
		Value: "root",
		Right: &BinTrie{
			Value: "Right 1",
			Right: &BinTrie{
				Value: "Right 2",
			},
			Left: &BinTrie{
				Value: "Right 1 Left",
			},
		},
		Left: &BinTrie{
			Value: "Left 1",
			Left: &BinTrie{
				Value: "Left 2",
				Left: &BinTrie{
					Value: "Left 3",
				},
				Right: &BinTrie{
					Value: "Left 2 Right",
				},
			},
		},
	}
	return root
}


func main() {
	root := createTestTrie()
	root.Dfs(func(binTrie *BinTrie) {
		fmt.Println(binTrie.Value)
	})
}
