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
	SingleBitMap struct {
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

	//BinaryTree struct
	BinaryTree struct {
		Root *BinaryTreeNode
	}

	BinaryTreeNode struct {
		V     interface{}
		Left  *BinaryTreeNode
		Right *BinaryTreeNode
	}

	node[T any] struct {
		pre  *node[T]
		v    T
		next *node[T]
	}
)
