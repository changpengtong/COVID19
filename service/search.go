package service

import "fmt"

type TrieNode struct {
	child map[rune] *TrieNode
	val interface{}
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{root: &TrieNode{child: make(map[rune]*TrieNode)}}
}

func (r *Trie) Insert(key string, val interface{}) {
	node :=r.root
	for _,c:=range key {
		if n,ok := node.child[c];!ok {
			newNode := &TrieNode{child:make(map[rune]*TrieNode)}
			node.child[c] = newNode
			node = newNode
		} else {
			node = n
		}
	}
	node.val = val
}

func (r *Trie) Get(key string) interface{} {
	node := r.root
	for _,c := range key {
		if n,ok:=node.child[c];!ok {
			return nil
		} else {
			node = n
		}
	}
	return node.val
}

func main() {
	r := newTrie()
	r.Insert("C","Covid,Covid19")
	r.Insert("Co","Corona,CoronaVirus,Corona")
	r.Insert("2","2020,2019")
	fmt.Println(r.Get("Co"))
}

