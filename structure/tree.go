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
	"math"
)

//NewBinaryTree 构造器
func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{V: v}
}

//PreOrder 前序遍历
func (bt *BinaryTree) PreOrder() []interface{} {
	res := make([]interface{}, 0)
	if bt != nil {
		res = append(res, bt.V)
		res = append(res, bt.Left.PreOrder()...)
		res = append(res, bt.Right.PreOrder()...)
	}

	return res
}

//MiddleOrder 中序遍历
func (bt *BinaryTree) MiddleOrder() []interface{} {
	res := make([]interface{}, 0)
	if bt != nil {
		res = append(res, bt.Left.MiddleOrder()...)
		res = append(res, bt.V)
		res = append(res, bt.Right.MiddleOrder()...)
	}

	return res
}

//PostOrder 后序遍历
func (bt *BinaryTree) PostOrder() []interface{} {
	res := make([]interface{}, 0)
	if bt != nil {
		res = append(res, bt.Left.PostOrder()...)
		res = append(res, bt.Right.PostOrder()...)
		res = append(res, bt.V)
	}

	return res
}

//IsBalanced check this tree is balanced
func (bt *BinaryTree) IsBalanced() bool {
	if recur(bt) == -1 {
		return false
	}
	return true
}

//High returns this tree high
func (bt *BinaryTree) High() int {
	if bt == nil {
		return 0
	}

	left := bt.Left.High()
	right := bt.Right.High()

	return int(math.Max(float64(left), float64(right))) + 1
}

func recur(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}
	if math.Abs(float64(left-right)) < 2 {
		//balance
		return int(math.Max(float64(left), float64(right))) + 1
	} else {
		return -1 //not balance
	}
}
