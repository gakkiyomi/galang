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

func (node *RBTreeNode[T]) GetValue() (data T) {
	if node == nil {
		return
	}
	return node.V
}

func (node *RBTreeNode[T]) GetThis() BBTreeNodeInterface[T] {
	return node
}

func (node *RBTreeNode[T]) GetLeft() BBTreeNodeInterface[T] {
	return node.Left
}

func (node *RBTreeNode[T]) GetRight() BBTreeNodeInterface[T] {
	return node.Right
}

func newNilNode[T comparable]() *RBTreeNode[T] {
	return &RBTreeNode[T]{Color: builtin.BLACK, IsEnd: true}
}

// 红黑树树构造器
func NewRBTree[T comparable](c builtin.Comparable[T]) *RBTree[T] {
	nilNode := newNilNode[T]()
	tree := &RBTree[T]{
		Root: &RBTreeNode[T]{
			Color: builtin.BLACK,
			Left:  nilNode,
			Right: nilNode,
			IsEnd: false,
		},
		compare: c,
	}
	return tree
}

func (self *RBTree[T]) Insert(v T) {
	insertInRB(self.compare, self.Root, v)
}

func (self *RBTree[T]) Delete(v T) {
	self.Root = deleteNodeInRB[T](self.compare, self.Root, v)
}

// Search search for the specified value in the tree
func (self *RBTree[T]) Search(v T) (data T, exist bool) {
	node, exist := search(self.compare, self.Root.GetThis(), v)
	if node == nil {
		return
	}
	return node.GetValue(), exist
}

// Max find the maximum or minimum value
func (self *RBTree[T]) Max() T {
	var max func(root *RBTreeNode[T]) T
	max = func(root *RBTreeNode[T]) T {
		if root.Right.IsEnd {
			return root.V
		}
		return max((root.Right))
	}
	return max((self.Root))
}

// Min find the maximum or minimum value
func (self *RBTree[T]) Min() T {
	var min func(root *RBTreeNode[T]) T
	min = func(root *RBTreeNode[T]) T {
		if root.Left.IsEnd {
			return root.V
		}
		return min((root.Left))
	}
	return min((self.Root))
}

func insertInRB[T comparable](c builtin.Comparable[T], root *RBTreeNode[T], v T) *RBTreeNode[T] {
	if nil == root || root.IsEnd {
		nilNode := newNilNode[T]()
		root = &RBTreeNode[T]{
			Color: builtin.BLACK,
			V:     v,
			Left:  nilNode,
			Right: nilNode,
			IsEnd: false,
		}
	}
	bfs(conditionPath[T])
	compares := c(v, root.V)
	if compares < 0 {
		root.Left = insertInRB(c, root.Left, v)
	} else if compares > 0 {
		root.Right = insertInRB(c, root.Right, v)
	} else {
		return root
	}
	// 插入后需要检查和修复红黑树的性质
	root = fixViolation(root)

	// 确保根节点是黑色
	if root.Color == builtin.RED {
		root.Color = builtin.BLACK
	}
	return root
}

func deleteNodeInRB[T comparable](c builtin.Comparable[T], root *RBTreeNode[T], v T) (res *RBTreeNode[T]) {
	return res
}

func fixViolation[T comparable](tree *RBTreeNode[T]) *RBTreeNode[T] {
	return nil
}

// 性质：如果一个节点是红色，那么他的子节点一定是黑色
func conditionRedNode[T comparable](node *RBTreeNode[T]) bool {
	if node.Color == builtin.BLACK {
		return true
	}
	if node.Left.Color == builtin.RED || node.Right.Color == builtin.RED {
		return false
	}
	return true
}

// 性质：每个节点的每条路径都包含相同数量的黑色节点
func conditionPath[T comparable](node *RBTreeNode[T]) (int, bool) {
	if nil == node || node.IsEnd {
		return 0, true
	}
	if node.Color == builtin.BLACK {
		blackLCount, ok1 := conditionPath(node.Left)
		blackRCount, ok2 := conditionPath(node.Right)
		if blackLCount == blackRCount && ok1 && ok2 {
			return blackLCount + 1, true
		} else {
			return 0, false
		}
	} else {
		blackLCount, ok1 := conditionPath(node.Left)
		blackRCount, ok2 := conditionPath(node.Right)
		if blackLCount == blackRCount && ok1 && ok2 {
			return blackLCount, true
		} else {
			return 0, false
		}
	}
}

func (self *RBTree[T]) BFS() bool {
	list := []*RBTreeNode[T]{self.Root}
	idx := 0
	end := len(list)
	for idx < end {
		n := list[idx]
		_, ok := conditionPath(n)
		if ok {
			return false
		}
		if !n.Left.IsEnd {
			list = append(list, n.Left)
			end++
		}
		if !n.Right.IsEnd {
			list = append(list, n.Right)
			end++
		}
		idx++
	}
	return true
}

// rotateRight 树节点右旋
func rotateRightRB[T comparable](root *RBTreeNode[T]) *RBTreeNode[T] {
	return nil
}

// rotateRight 树节点左旋
func rotateLeftRB[T comparable](root *RBTreeNode[T]) *RBTreeNode[T] {
	return nil
}
