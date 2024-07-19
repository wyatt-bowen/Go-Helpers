package trie

import "strings"

// This is an open ended trie intended for the word search. This means that words in here are all valid, they don't have an ending tag. In T -> R -> I -> E, t, tr, tri, trie are all valid words

type Trie struct {
	start  bool
	letter rune
	next   []Trie
}

func NewTrie() *Trie {
	return &Trie{true, '.', make([]Trie, 0)}
}

func NewTrieNode(c rune) *Trie {
	return &Trie{false, c, make([]Trie, 0)}
}

func (t *Trie) AddWord(word string) {
	if !t.start {
		return
	}
	word = strings.ToUpper(word)
	t.addWordHelper(word)
}

func (t *Trie) addWordHelper(word string) {
	if len(word) == 0 {
		return
	}
	nextLetter := rune(word[0])
	restOfWord := string(word[1:])
	if loc, contains := containsLetter(t.next, nextLetter); contains {
		t.next[loc].addWordHelper(restOfWord)
	} else {
		t.next = append(t.next, *NewTrieNode(nextLetter))
		t.next[len(t.next)-1].addWordHelper(restOfWord)
	}
}

func containsLetter(s []Trie, letter rune) (int, bool) {
	for i, v := range s {
		if v.letter == letter {
			return i, true
		}
	}
	return 0, false
}

func (t *Trie) ContainsWord(word string) bool {
	word = strings.ToUpper(word)
	return t.containsWordHelper(word)
}

func (t *Trie) containsWordHelper(word string) bool {
	if len(word) == 0 {
		panic("len0")
	} else if len(word) == 1 {
		_, result := containsLetter(t.next, rune(word[0]))
		return result
	} else {
		nextLetter := rune(word[0])
		restOfWord := string(word[1:])
		if i, contains := containsLetter(t.next, nextLetter); contains {
			return t.next[i].containsWordHelper(restOfWord)
		} else {
			return false
		}
	}
}

func NewTrieFromWord(word string) *Trie {
	result := NewTrie()
	result.AddWord(word)
	return result
}
