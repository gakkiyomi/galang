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
	"reflect"

	"github.com/gakkiyomi/galang/builtin"
)

func (node *BinaryTreeNode[T]) GetValue() (data T) {
	if node == nil {
		return
	}
	return node.V
}

func (node *BinaryTreeNode[T]) GetThis() BBTreeNodeInterface[T] {
	return node
}

func (node *BinaryTreeNode[T]) GetLeft() BBTreeNodeInterface[T] {
	return node.Left
}

func (node *BinaryTreeNode[T]) GetRight() BBTreeNodeInterface[T] {
	return node.Right
}

// AVL树构造器
func NewAVLTree[T comparable](c builtin.Comparable[T]) *AVLTree[T] {
	tree := &AVLTree[T]{
		BinaryTree: &BinaryTree[T]{},
		compare:    c,
	}
	return tree
}

func (self *AVLTree[T]) Insert(v T) {
	insertNode(self, v)
}

func (self *AVLTree[T]) Delete(v T) {
	self.Root = deleteNode[T](self.compare, self.Root, v)
}

// Search search for the specified value in the tree
func (self *AVLTree[T]) Search(v T) (data T, exist bool) {
	node, exist := search(self.compare, self.Root.GetThis(), v)
	if node == nil {
		return
	}
	return node.GetValue(), exist
}

// Max find the maximum or minimum value
func (self *AVLTree[T]) Max() T {
	return findOp(self.Root, builtin.RIGHT)
}

// Min find the maximum or minimum value
func (self *AVLTree[T]) Min() T {
	return findOp(self.Root, builtin.LEFT)
}

func insert[T comparable](c builtin.Comparable[T], root *BinaryTreeNode[T], v T) *BinaryTreeNode[T] {
	if nil == root {
		return &BinaryTreeNode[T]{V: v, Height: 0}
	}
	compareRes := c(v, root.V)
	if compareRes < 0 {
		root.Left = insert(c, root.Left, v)
	} else if compareRes > 0 {
		root.Right = insert(c, root.Right, v)
	} else {
		return root
	}
	return root
}

func updateHeight[T comparable](node *BinaryTreeNode[T]) {
	if node.Left != nil {
		updateHeight(node.Left)
	}
	if node.Right != nil {
		updateHeight(node.Right)
	}
	node.Height = 1 + int(math.Max(getHeight(node.Left), getHeight(node.Right)))
}

func deleteNode[T comparable](c builtin.Comparable[T], root *BinaryTreeNode[T], v T) *BinaryTreeNode[T] {
	if root == nil {
		return root
	}
	compareRes := c(v, root.V)
	if compareRes < 0 {
		root.Left = deleteNode[T](c, root.Left, v)
	} else if compareRes > 0 {
		root.Right = deleteNode[T](c, root.Right, v)
	} else {
		if root.Left == nil || root.Right == nil {
			var temp *BinaryTreeNode[T]
			if root.Left != nil {
				temp = root.Left
			} else {
				temp = root.Right
			}
			if temp == nil {
				//don't need to action
				temp = root
				root = nil
			} else {
				//temp => root
				*root = *temp
			}
			temp = nil
		} else {
			root.V = findOp[T](root.Right, builtin.LEFT)
			root.Right = deleteNode[T](c, root.Right, root.V)
		}
		if root == nil {
			return root
		}
		updateHeight[T](root)
		balance := GetBalance(root)
		leftBalance := GetBalance(root.Left)
		rightBalance := GetBalance(root.Right)
		// Left Left Case
		if balance > 1 && leftBalance >= 0 {
			return rotateRight(root)

		}

		// Right Right Case
		if balance < -1 && rightBalance <= 0 {
			return rotateLeft(root)

		}

		// Left Right Case
		if balance > 1 && leftBalance < 0 {
			root.Left = rotateLeft(root.Left)
			return rotateRight(root)

		}

		// Right Left Case
		if balance < -1 && rightBalance > 0 {
			root.Right = rotateRight(root.Right)
			return rotateLeft(root)
		}
	}
	return root
}

func insertNode[T comparable](tree *AVLTree[T], v T) {
	tree.Root = insert[T](tree.compare, tree.Root, v)
	//update tree height
	updateHeight[T](tree.Root)
	balance := GetBalance[T](tree.Root)
	if balance == 0 {
		return
	}
	root := tree.Root
	// Left Left Case
	if balance > 1 && tree.compare(v, root.Left.V) < 0 {
		tree.Root = rotateRight(tree.Root)
		return
	}

	// Right Right Case
	if balance < -1 && tree.compare(v, root.Right.V) > 0 {
		tree.Root = rotateLeft(root)
		return
	}

	// Left Right Case
	if balance > 1 && tree.compare(v, root.Left.V) > 0 {
		tree.Root.Left = rotateLeft(root.Left)
		root = rotateRight(root)
		return
	}

	// Right Left Case
	if balance < -1 && tree.compare(v, root.Right.V) < 0 {
		tree.Root.Right = rotateRight(root.Right)
		root = rotateLeft(root)
		return
	}

}

// 通过接口实现函数复用
func search[T comparable](c builtin.Comparable[T], root BBTreeNodeInterface[T], v T) (BBTreeNodeInterface[T], bool) {
	if root == nil || reflect.ValueOf(root).IsNil() {
		return nil, false
	}
	compareRes := c(v, root.GetValue())
	if compareRes < 0 {
		return search[T](c, root.GetLeft(), v)
	} else if compareRes > 0 {
		return search[T](c, root.GetRight(), v)
	} else {
		return root, true
	}
}

// findOp Recursively find the maximum or minimum value
func findOp[T comparable](root *BinaryTreeNode[T], dirction builtin.Direction) (data T) {
	var leaf *BinaryTreeNode[T]
	if dirction == builtin.LEFT {
		leaf = root.Left
	} else {
		leaf = root.Right
	}
	if leaf == nil {
		data = root.V
		return
	}
	data = findOp(leaf, dirction)
	return
}

// getHeight get height of current tree node
func getHeight[T comparable](node *BinaryTreeNode[T]) float64 {
	if node == nil {
		return 0
	}
	return float64(node.Height)
}

// GetBalance get a factors to represent tree balance data
func GetBalance[T comparable](root *BinaryTreeNode[T]) int {
	if root == nil {
		return 0
	}
	return int(getHeight(root.Left) - getHeight(root.Right))
}

// rotateRight 树节点右旋
func rotateRight[T comparable](root *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	newlyRoot := root.Left
	root.Left = newlyRoot.Right
	newlyRoot.Right = root

	updateHeight(newlyRoot)
	updateHeight(root)
	return newlyRoot
}

// rotateRight 树节点左旋
func rotateLeft[T comparable](root *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	newlyRoot := root.Right
	root.Right = newlyRoot.Left
	newlyRoot.Left = root

	updateHeight(newlyRoot)
	updateHeight(root)
	return newlyRoot
}
