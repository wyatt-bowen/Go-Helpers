package trie

import "strings"

// This is an open ended trie intended for the word search. This means that words in here are all valid, they don't have an ending tag. In T -> R -> I -> E, t, tr, tri, trie are all valid words

type OpenTrie struct {
	start  bool
	letter rune
	next   []OpenTrie
}

func NewOpenTrie() *OpenTrie {
	return &OpenTrie{true, '.', make([]OpenTrie, 0)}
}

func NewOpenTrieNode(c rune) *OpenTrie {
	return &OpenTrie{false, c, make([]OpenTrie, 0)}
}

func (t *OpenTrie) AddWord(word string) {
	if !t.start {
		return
	}
	word = strings.ToUpper(word)
	t.addWordHelper(word)
}

func (t *OpenTrie) addWordHelper(word string) {
	if len(word) == 0 {
		return
	}
	nextLetter := rune(word[0])
	restOfWord := string(word[1:])
	if loc, contains := containsLetterOpen(t.next, nextLetter); contains {
		t.next[loc].addWordHelper(restOfWord)
	} else {
		t.next = append(t.next, *NewOpenTrieNode(nextLetter))
		t.next[len(t.next)-1].addWordHelper(restOfWord)
	}
}

func containsLetterOpen(s []OpenTrie, letter rune) (int, bool) {
	for i, v := range s {
		if v.letter == letter {
			return i, true
		}
	}
	return 0, false
}

func (t *OpenTrie) ContainsWord(word string) bool {
	word = strings.ToUpper(word)
	return t.containsWordHelper(word)
}

func (t *OpenTrie) containsWordHelper(word string) bool {
	if len(word) == 0 {
		panic("len0")
	} else if len(word) == 1 {
		_, result := containsLetterOpen(t.next, rune(word[0]))
		return result
	} else {
		nextLetter := rune(word[0])
		restOfWord := string(word[1:])
		if i, contains := containsLetterOpen(t.next, nextLetter); contains {
			return t.next[i].containsWordHelper(restOfWord)
		} else {
			return false
		}
	}
}
