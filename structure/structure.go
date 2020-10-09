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

type (
	//Stack interface
	Stack interface {
		Push(interface{})
		Pop() interface{}
		Peek() interface{}
		Len() int
		IsEmpty() bool
	}

	//LinkedStack is use Linked list achieve stack
	LinkedStack struct {
		head *node
		len  int
	}

	//Queue struct
	Queue struct {
		head *node
		tail *node
		len  int
	}

	//BinaryTree struct
	BinaryTreeNode struct {
		V     interface{}
		Left  *BinaryTreeNode
		Right *BinaryTreeNode
	}

	BinaryTree struct {
		Root *BinaryTreeNode
	}

	node struct {
		pre  *node
		v    interface{}
		next *node
	}
)
