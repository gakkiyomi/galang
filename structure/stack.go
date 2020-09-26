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

//NewStack returns a Stack
func NewStack() Stack {
	return &LinkedStack{nil, 0}
}

//Len returns a stack length
func (stack *LinkedStack) Len() int {
	return stack.len
}

//Push a value into stack
func (stack *LinkedStack) Push(v interface{}) {
	n := &node{nil, v, stack.head}
	stack.head = n
	stack.len++
}

//Pop a vlaue out of stack
func (stack *LinkedStack) Pop() interface{} {
	if stack.len == 0 {
		return nil
	}
	res := stack.head
	stack.head = res.next
	stack.len--
	return res.v
}

//Peek returns the head of stack
func (stack *LinkedStack) Peek() interface{} {
	if stack.len == 0 {
		return nil
	}
	return stack.head.v
}

//IsEmpty 栈是否为空
func (stack *LinkedStack) IsEmpty() bool {
	return stack.Len() == 0
}
