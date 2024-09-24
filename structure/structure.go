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
	"github.com/gakkiyomi/galang/builtin"
)

type (
	BitMap struct {
		bitmaps []uint64
		len     int
	}
	//Stack interface
	Stack[T any] interface {
		Push(T)
		Pop() T
		Peek() T
		Len() int
		IsEmpty() bool
	}

	//RingBuffer
	RingBuffer[T any] interface {
		Put(T) bool
		Get() T
		GetAll() []T
		Len() int
		IsFull() bool
		IsEmpty() bool
	}

	ArrayRingBuffer[T any] struct {
		buffer     []T
		cap        int
		len        int
		wirteIndex int
		readIndex  int
	}

	//LinkedStack is use Linked list achieve stack
	LinkedStack[T any] struct {
		head *node[T]
		len  int
	}

	//Queue struct
	Queue[T any] struct {
		head *node[T]
		tail *node[T]
		len  int
	}

	//Heap struct
	Heap struct {
		Items []int
	}

	//AVLTree the ordinary balanced binary tree
	AVLTree[T comparable] struct {
		*BinaryTree[T]
		// compare function
		compare builtin.Comparable[T]
	}

	//RBTree
	RBTree[T comparable] struct {
		Root *RBTreeNode[T]
		// compare function
		compare builtin.Comparable[T]
	}

	//BalancedBinaryTree interface Definition an BalancedBinaryTree need implement methods
	BalancedBinaryTree[T comparable] interface {
		Insert(v T)
		Delete(v T)
		Search(v T) T
		Max() T
		Min() T
		RotateRight(BBTreeNodeInterface[T]) BBTreeNodeInterface[T]
		RotateLeft(BBTreeNodeInterface[T]) BBTreeNodeInterface[T]
		SetLeft(BBTreeNodeInterface[T])
		SetRight(BBTreeNodeInterface[T])
	}

	//BinaryTree struct
	BinaryTree[T comparable] struct {
		Root *BinaryTreeNode[T]
	}

	BinaryTreeNode[T comparable] struct {
		V      T
		Height int
		Left   *BinaryTreeNode[T]
		Right  *BinaryTreeNode[T]
	}

	BBTreeNodeInterface[T comparable] interface {
		GetThis() BBTreeNodeInterface[T]
		GetValue() T
		GetLeft() BBTreeNodeInterface[T]
		GetRight() BBTreeNodeInterface[T]
	}

	RBTreeNode[T comparable] struct {
		V      T
		Color  builtin.Color
		Left   *RBTreeNode[T]
		Right  *RBTreeNode[T]
		Parent *RBTreeNode[T]
	}

	//Tree common tree struct
	Tree[T comparable] struct {
		Root *TreeNode[T]
	}

	//TreeNode common tree node struct
	TreeNode[T comparable] struct {
		V      T
		Childs []*TreeNode[T]
		Extra  any
	}

	node[T any] struct {
		pre  *node[T]
		v    T
		next *node[T]
	}
)
