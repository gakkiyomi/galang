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

// Tire trie strcut
type Trie Tree[rune]

func NewTrie() (res *Trie) {
	res = &Trie{
		Root: &TreeNode[rune]{
			V:      0,
			Childs: make([]*TreeNode[rune], 0),
			Extra:  false,
		},
	}
	return
}

// Insert insert a string to trie.
func (t *Trie) Insert(word string) {
	currentNode := t.Root
	for _, char := range word {
		found := false
		// 检查当前节点的子节点是否包含当前字符
		for _, child := range currentNode.Childs {
			if child.V == char {
				currentNode = child
				found = true
				break
			}
		}
		// 如果没有找到当前字符的子节点，则创建一个新的子节点并移动到该节点
		if !found {
			newNode := &TreeNode[rune]{V: char, Childs: make([]*TreeNode[rune], 0)}
			currentNode.Childs = append(currentNode.Childs, newNode)
			currentNode = newNode
		}
	}
	//设置当前节点为某个字符串的end节点
	currentNode.Extra = true
}

// Insert delete a string from trie.
func (t *Trie) Delete(word string) {
	t.deleteReur(word, t.Root, 0)
}

func (t *Trie) deleteReur(word string, currentNode *TreeNode[rune], index int) {
	if index == len(word) { //如果到了最后一个字符 递归退出
		return
	}
	char := rune(word[index])

	var childNode *TreeNode[rune]
	for _, child := range currentNode.Childs {
		if child.V == char {
			childNode = child
			break
		}
	}
	if childNode == nil { //退出
		return
	}
	t.deleteReur(word, childNode, index+1)
	if len(childNode.Childs) == 0 {
		for i, child := range currentNode.Childs {
			if child == childNode {
				currentNode.Childs = append(currentNode.Childs[:i], currentNode.Childs[i+1:]...)
				break
			}
		}
	}
}

func (t *Trie) Search(word string) (exist bool) {
	currentNode := t.Root
	for _, char := range word {
		exist = false
		// 检查当前节点的子节点是否包含当前字符
		for _, child := range currentNode.Childs {
			if child.V == char {
				currentNode = child
				exist = true
				break
			}
		}
		if !exist {
			return
		}
	}
	return
}

// PrefixSearch get all strings with a common prefix.
func (t *Trie) PrefixSearch(word string) []string {
	currentNode := t.Root
	for _, char := range word {
		found := false
		// 检查当前节点的子节点是否包含当前字符
		for _, child := range currentNode.Childs {
			if child.V == char {
				currentNode = child
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	// 找到具有相同前缀的所有字符串
	var result []string
	dfs(currentNode, word, &result)
	return result
}

func dfs(node *TreeNode[rune], currentPrefix string, result *[]string) {
	if node.Extra == true {
		*result = append(*result, currentPrefix)
	}
	if len(node.Childs) == 0 {
		return
	}
	for _, child := range node.Childs {
		dfs(child, currentPrefix+string(child.V), result)
	}
}
