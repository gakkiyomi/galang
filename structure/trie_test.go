// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieInsertSearchDelete(t *testing.T) {
	// 创建一个新的 Trie
	trie := NewTrie()

	// 插入一些单词到 Trie 中
	words := []string{"apple", "app", "apricot", "bat", "batman"}
	for _, word := range words {
		trie.Insert(word)
	}

	// 在 Trie 中搜索单词
	assert.True(t, trie.Search("apple"), "Expected 'apple' to be in Trie")
	assert.True(t, trie.Search("app"), "Expected 'app' to be in Trie")
	assert.True(t, trie.Search("apricot"), "Expected 'apricot' to be in Trie")
	assert.False(t, trie.Search("banana"), "Expected 'banana' not to be in Trie")
	assert.True(t, trie.Search("batman"), "Expected 'batman' to be in Trie")
	assert.False(t, trie.Search("batwoman"), "Expected 'batwoman' not to be in Trie")

	// 删除一个单词
	trie.Delete("apple")

	// 重新检查单词是否存在
	assert.False(t, trie.Search("apple"), "Expected 'apple' not to be in Trie")
}

func TestTriePrefixSearch(t *testing.T) {
	// 创建一个新的 Trie
	trie := NewTrie()

	// 插入一些单词到 Trie 中
	words := []string{"apple", "app", "apricot", "bat", "batman"}
	for _, word := range words {
		trie.Insert(word)
	}

	// PrefixSearch
	assert.ElementsMatch(t, trie.PrefixSearch("ap"), []string{"app", "apple", "apricot"}, "Expected ['app', 'apricot'] for prefix 'ap'")
	assert.ElementsMatch(t, trie.PrefixSearch("bat"), []string{"bat", "batman"}, "Expected ['bat', 'batman'] for prefix 'bat'")
	assert.ElementsMatch(t, trie.PrefixSearch("fang"), nil)
}
