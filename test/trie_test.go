package test

import (
	"fmt"
	"testing"

	"github.com/wyatt-bowen/go-helpers/collections/trie"
	"github.com/wyatt-bowen/go-helpers/slicehelpers"
)

func TestOpenTrie(t *testing.T) {
	testTrie := trie.NewOpenTrie()
	testTrie.AddWord("testing")
	testTrie.AddWord("trie")
	if !testTrie.ContainsWord("test") {
		t.Fail()
	}
	if !testTrie.ContainsWord("trie") {
		t.Fail()
	}
	if testTrie.ContainsWord("tree") {
		t.Fail()
	}
}

func TestTrie(t *testing.T) {
	testTrie := trie.NewTrie()
	testTrie.AddWord("testing")
	testTrie.AddWord("trie")
	if testTrie.ContainsWord("test") {
		fmt.Println("testTrie incorrectly claims to contain test")
		t.Fail()
	}
	if !testTrie.ContainsWord("trie") {
		t.Fail()
	}
	if testTrie.ContainsWord("tree") {
		t.Fail()
	}
	if !testTrie.ContainsWord("testing") {
		t.Fail()
	}
}

func TestTrieGetWords(t *testing.T) {
	testTrie := trie.NewTrie()
	testTrie.AddWord("testing")
	testTrie.AddWord("trie")
	words := testTrie.GetAllWords()

	fmt.Println(slicehelpers.ToString(words))
}
