package main

import "fmt"

type Trie struct {
	Value    string
	Children []*Trie
}

func (t *Trie) Dfs(f func(trie *Trie)) {
	if t == nil {
		return
	}
	f(t)
	for _, child := range t.Children {
		child.Dfs(f)
	}
}

func createTestTrie() *Trie {
	return &Trie{
		Value: "root",
		Children: []*Trie{
			{Value: "root ch 1"},
			{Value: "root ch 2"},
			{Value: "root ch 3", Children: []*Trie{
				{Value: "ch 31"},
				{Value: "ch 32", Children: []*Trie{
					{Value: "ch 321", Children: []*Trie{
						{Value: "ch 3211"},
						{Value: "ch 3212"},
						{Value: "ch 3213"},
					}},
				}},
			}},
			{Value: "root ch 4"},
		},
	}
}

func main() {
	root := createTestTrie()
	root.Dfs(func(trie *Trie) {
		fmt.Println(trie.Value)
	})
}
