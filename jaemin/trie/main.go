package main

import (
	"fmt"
	"sync"
)

type node struct {
	children map[string]*node
	isLast   bool
}

type Trie struct {
	root *node
	*sync.RWMutex
}

func newNode() *node {
	return &node{
		children: make(map[string]*node),
		isLast:   false,
	}
}

func NewTrie() Trie {
	rootNode := &node{
		children: make(map[string]*node),
		isLast:   true,
	}

	return Trie{
		root:    rootNode,
		RWMutex: &sync.RWMutex{},
	}
}

func (t *Trie) Insert(str string) {
	t.Lock()
	defer t.Unlock()

	head := t.root
	t.root.isLast = false

	for index := range str {

		char := string(str[index])
		if _, exists := head.children[char]; !exists {
			head.children[char] = newNode()
		}

		head = head.children[char]
	}

	head.isLast = true
}

func (t Trie) Search(str string) bool {
	t.RLock()
	defer t.RUnlock()

	head := t.root

	for index := range str {
		char := string(str[index])
		if head.children[char] == nil {
			return false
		}

		head = head.children[char]
	}

	return head.isLast
}

func (t *Trie) Delete(str string) {
	t.Lock()
	defer t.Unlock()

	delete(t.root, str, 0)
}

func delete(head *node, str string, depth int) (isDeleted, isEmpty bool) {
	if depth == len(str) {
		if head.isLast {
			head.isLast = false
		}

		if len(head.children) == 0 {
			return true, true
		}

		return true, false
	}

	char := string(str[depth])
	if child := head.children[char]; child != nil {
		isDeleted, isEmpty = delete(child, str, depth+1)
		if isDeleted && isEmpty {
			head.children[char] = nil
		}

		isEmpty = len(head.children) == 0
	}

	return
}

func main() {
	strings := []string{"hello", "hi", "here", "fjdsngfjasd"}

	trie := NewTrie()
	for _, str := range strings {
		trie.Insert(str)
	}

	fmt.Println(trie.Search("h")) // false
	fmt.Println(trie.Search(""))  // false
	fmt.Println("---------------")

	fmt.Println(trie.Search("hello"))       // true
	fmt.Println(trie.Search("hi"))          // true
	fmt.Println(trie.Search("here"))        // true
	fmt.Println(trie.Search("fjdsngfjasd")) // true
	fmt.Println("---------------")

	trie.Delete("hello")
	fmt.Println(trie.Search("hello"))       // false
	fmt.Println(trie.Search("hi"))          // true
	fmt.Println(trie.Search("here"))        // true
	fmt.Println(trie.Search("fjdsngfjasd")) // true
	fmt.Println("---------------")

	trie.Delete("fjdsngfjasd")
	fmt.Println(trie.Search("hello"))       // false
	fmt.Println(trie.Search("hi"))          // true
	fmt.Println(trie.Search("here"))        // true
	fmt.Println(trie.Search("fjdsngfjasd")) // false
	fmt.Println("---------------")

	trie.Delete("hi")
	fmt.Println(trie.Search("hello"))       // false
	fmt.Println(trie.Search("hi"))          // false
	fmt.Println(trie.Search("here"))        // true
	fmt.Println(trie.Search("fjdsngfjasd")) // false
	fmt.Println("---------------")

	trie.Delete("here")
	fmt.Println(trie.Search("hello"))       // false
	fmt.Println(trie.Search("hi"))          // false
	fmt.Println(trie.Search("here"))        // false
	fmt.Println(trie.Search("fjdsngfjasd")) // false
	fmt.Println("---------------")

	trie.Delete("here")
	fmt.Println(trie.Search("hello"))       // false
	fmt.Println(trie.Search("hi"))          // false
	fmt.Println(trie.Search("here"))        // false
	fmt.Println(trie.Search("fjdsngfjasd")) // false
	fmt.Println("---------------")

	fmt.Println(trie.Search("hello"))       // false
	fmt.Println(trie.Search("hi"))          // false
	fmt.Println(trie.Search("here"))        // false
	fmt.Println(trie.Search("fjdsngfjasd")) // false
}
