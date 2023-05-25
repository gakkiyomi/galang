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

// NewBinaryTree 构造器
func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{
		Root: &BinaryTreeNode{
			V: v,
		}}
}

func AddNode(v interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		V: v,
	}
}

// PreOrder 前序遍历
func (bt *BinaryTree) PreOrder() []interface{} {
	return preOrder(bt.Root)
}

func preOrder(root *BinaryTreeNode) []interface{} {
	res := make([]interface{}, 0)
	if root != nil {
		res = append(res, root.V)
		res = append(res, preOrder(root.Left)...)
		res = append(res, preOrder(root.Right)...)
	}
	return res
}

// MiddleOrder 中序遍历
func (bt *BinaryTree) MiddleOrder() []interface{} {
	return middleOrder(bt.Root)
}

func middleOrder(root *BinaryTreeNode) []interface{} {
	res := make([]interface{}, 0)
	if root != nil {
		res = append(res, middleOrder(root.Left)...)
		res = append(res, root.V)
		res = append(res, middleOrder(root.Right)...)
	}

	return res
}

// PostOrder 后序遍历
func (bt *BinaryTree) PostOrder() []interface{} {
	return postOrder(bt.Root)
}

func postOrder(root *BinaryTreeNode) []interface{} {
	res := make([]interface{}, 0)
	if root != nil {
		res = append(res, postOrder(root.Left)...)
		res = append(res, postOrder(root.Right)...)
		res = append(res, root.V)
	}

	return res
}

// BFS 层次遍历
func (bt *BinaryTree) BFS() []interface{} {
	res := make([]interface{}, 0)
	if bt != nil {
		nodes := []*BinaryTreeNode{bt.Root}
		for len(nodes) > 0 {
			currentNode := nodes[0]
			nodes = nodes[1:]
			res = append(res, currentNode.V)
			if currentNode.Left != nil {
				nodes = append(nodes, currentNode.Left)
			}
			if currentNode.Right != nil {
				nodes = append(nodes, currentNode.Right)
			}
		}
	}
	return res
}

// IsBalanced check this tree is balanced
func (bt *BinaryTree) IsBalanced() bool {
	if recur(bt.Root) == -1 {
		return false
	}
	return true
}

// High returns this tree high
func (bt *BinaryTree) High() int {
	return high(bt.Root)
}

func high(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	left := high(root.Left)
	right := high(root.Right)

	return int(math.Max(float64(left), float64(right))) + 1
}

// Size returns this tree node size
func (bt *BinaryTree) Size() int {
	return size(bt.Root)
}

func size(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	left := size(root.Left)
	right := size(root.Right)

	return left + right + 1
}

func recur(root *BinaryTreeNode) int {
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
	}
	return -1 //not balance

}
