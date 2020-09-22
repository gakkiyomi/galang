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

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
}

type LinkedStack struct {
	head *node
	len  int
}

type node struct {
	v    interface{}
	next *node
}

func NewStack() Stack {
	return &LinkedStack{nil, 0}
}

func (stack *LinkedStack) Len() int {
	return stack.len
}

func (stack *LinkedStack) Push(v interface{}) {
	n := &node{v, stack.head}
	stack.head = n
	stack.len++
}

func (stack *LinkedStack) Pop() interface{} {
	if stack.len == 0 {
		return nil
	}
	res := stack.head
	stack.head = res.next
	stack.len--
	return res.v
}

func (stack *LinkedStack) Peek() interface{} {
	if stack.len == 0 {
		return nil
	}
	return stack.head.v
}
