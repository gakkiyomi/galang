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
	"fmt"

	"github.com/gakkiyomi/galang/builtin"
	"github.com/gakkiyomi/galang/utils"
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

// 红黑树树构造器
func NewRBTree[T comparable](c builtin.Comparable[T]) *RBTree[T] {
	tree := &RBTree[T]{
		compare: c,
	}
	return tree
}

func (self *RBTree[T]) Insert(v T) {
	if self.Root == nil {
		self.Root = &RBTreeNode[T]{
			Color:  builtin.BLACK,
			Left:   nil,
			Right:  nil,
			Parent: nil,
			V:      v,
		}
		return
	}
	var insertedNode *RBTreeNode[T]
	currentNode := self.Root
outLoop:
	for {
		if currentNode == nil {
			break
		}
		compare := self.compare(v, currentNode.V)
		switch {
		case compare == 0:
			currentNode.V = v
			return
		case compare > 0: //插入右边
			if currentNode.Right == nil {
				insertedNode = &RBTreeNode[T]{
					V:      v,
					Color:  builtin.RED,
					Parent: currentNode,
				}
				currentNode.Right = insertedNode
				break outLoop
			}
			currentNode = currentNode.Right
		case compare < 0: //插入左边
			if currentNode.Left == nil {
				insertedNode = &RBTreeNode[T]{
					V:      v,
					Color:  builtin.RED,
					Parent: currentNode,
				}
				currentNode.Left = insertedNode
				break outLoop
			}
			currentNode = currentNode.Left
		}
	}
	self.rebalanced(insertedNode)
}

func (self *RBTree[T]) Delete(v T) {
	var deleteNode *RBTreeNode[T]
	currentNode := self.Root
outLoop:
	for currentNode != nil {
		compare := self.compare(v, currentNode.V)
		switch {
		case compare == 0:
			//delete this node
			deleteNode = currentNode
			break outLoop
		case compare > 0:
			currentNode = currentNode.Right
		case compare < 0:
			currentNode = currentNode.Left
		}
	}
	self.rebalancedWhenDeleted(deleteNode)
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
		if root.Right == nil {
			return root.V
		}
		return max((root.Right))
	}
	return max((self.Root))
}

func (self *RBTreeNode[T]) grandparent() (grandparent *RBTreeNode[T]) {
	if self == nil || self.Parent == nil || self.Parent.Parent == nil {
		return
	}
	grandparent = self.Parent.Parent
	return
}

// sibling 获取相邻节点
func (self *RBTreeNode[T]) sibling() *RBTreeNode[T] {
	if self == nil || self.Parent == nil {
		return nil
	}
	if self.Parent.Left == self {
		return self.Parent.Right
	}
	return self.Parent.Left
}

func (self *RBTreeNode[T]) uncle() (uncle *RBTreeNode[T]) {
	if self == nil || self.Parent == nil {
		return nil
	}
	uncle = self.Parent.sibling()
	return
}

// Min find the maximum or minimum value
func (self *RBTree[T]) Min() T {
	var min func(root *RBTreeNode[T]) T
	min = func(root *RBTreeNode[T]) T {
		if root.Left == nil {
			return root.V
		}
		return min((root.Left))
	}
	return min((self.Root))
}

func (tree *RBTree[T]) rebalancedWhenDeleted(node *RBTreeNode[T]) {
	if node == nil {
		return
	}
	var child *RBTreeNode[T]
	if node.Left != nil && node.Right != nil {
		prev := node.maximumNode() //如果左右节点都不为空，找到最大前驱节点
		node.V = prev.V
		node = prev //node指向最大前驱节点（等价于交换了node和prev的值之后删除prev就可以了）删除node => 删除prev
	}
	if node.Left != nil || node.Right != nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if nodeColor(node) == builtin.BLACK {
			node.Color = nodeColor(child) //将孩子节点的颜色赋给node
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.Color = builtin.BLACK
		}
	}
}

func (tree *RBTree[T]) deleteCase1(node *RBTreeNode[T]) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *RBTree[T]) deleteCase2(node *RBTreeNode[T]) {
	sibling := node.sibling()
	if nodeColor(sibling) == builtin.RED {
		node.Parent.Color = builtin.RED
		sibling.Color = builtin.BLACK
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *RBTree[T]) deleteCase3(node *RBTreeNode[T]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == builtin.BLACK &&
		nodeColor(sibling) == builtin.BLACK &&
		nodeColor(sibling.Left) == builtin.BLACK &&
		nodeColor(sibling.Right) == builtin.BLACK {
		sibling.Color = builtin.RED
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *RBTree[T]) deleteCase4(node *RBTreeNode[T]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == builtin.RED &&
		nodeColor(sibling) == builtin.BLACK &&
		nodeColor(sibling.Left) == builtin.BLACK &&
		nodeColor(sibling.Right) == builtin.BLACK {
		sibling.Color = builtin.RED
		node.Parent.Color = builtin.BLACK
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *RBTree[T]) deleteCase5(node *RBTreeNode[T]) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == builtin.BLACK &&
		nodeColor(sibling.Left) == builtin.RED &&
		nodeColor(sibling.Right) == builtin.BLACK {
		sibling.Color = builtin.RED
		sibling.Left.Color = builtin.BLACK
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == builtin.BLACK &&
		nodeColor(sibling.Right) == builtin.RED &&
		nodeColor(sibling.Left) == builtin.BLACK {
		sibling.Color = builtin.RED
		sibling.Right.Color = builtin.BLACK
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *RBTree[T]) deleteCase6(node *RBTreeNode[T]) {
	sibling := node.sibling()
	sibling.Color = nodeColor(node.Parent)
	node.Parent.Color = builtin.BLACK
	if node == node.Parent.Left && nodeColor(sibling.Right) == builtin.RED {
		sibling.Right.Color = builtin.BLACK
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == builtin.RED {
		sibling.Left.Color = builtin.BLACK
		tree.rotateRight(node.Parent)
	}
}

func (node *RBTreeNode[T]) maximumNode() *RBTreeNode[T] {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *RBTree[T]) rebalanced(node *RBTreeNode[T]) {
	tree.insertCase1(node)
}

func (tree *RBTree[T]) insertCase1(node *RBTreeNode[T]) {
	if node.Parent == nil { //Root
		node.Color = builtin.BLACK
		return
	}
	tree.insertCase2(node)
}

func (tree *RBTree[T]) insertCase2(node *RBTreeNode[T]) {
	if node.Parent.Color == builtin.BLACK { //如果父节点是黑色，直接跳过
		return
	}
	tree.insertCase3(node)
}

func (tree *RBTree[T]) insertCase3(node *RBTreeNode[T]) {
	uncle := node.uncle()
	if nodeColor(uncle) == builtin.RED { //如果叔叔节点是红色，那么更改颜色
		uncle.Color = builtin.BLACK
		node.Parent.Color = builtin.BLACK
		grandparent := node.grandparent()
		grandparent.Color = builtin.RED
		tree.insertCase1(grandparent) //祖父节点变红，递归处理
	} else {
		//叔叔节点为黑色，进行旋转处理
		tree.insertCase4(node)
	}
}

func (tree *RBTree[T]) insertCase4(node *RBTreeNode[T]) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left { //如果是之字型也就是LR
		//对父节点进行左旋
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right { //如果是之字型也就是RL
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)

}
func (tree *RBTree[T]) insertCase5(node *RBTreeNode[T]) {
	node.Parent.Color = builtin.BLACK
	grandparent := node.grandparent()
	grandparent.Color = builtin.RED
	if node == node.Parent.Left && node.Parent == grandparent.Left { //LL
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right { //RR
		tree.rotateLeft(grandparent)
	}
}

func nodeColor[T comparable](node *RBTreeNode[T]) builtin.Color {
	if node == nil {
		return builtin.BLACK
	}
	return node.Color
}

// 性质：如果一个节点是红色，那么他的子节点一定是黑色
func conditionRedNode[T comparable](node *RBTreeNode[T]) bool {
	if nodeColor(node) == builtin.BLACK {
		return true
	}
	if nodeColor(node.Left) == builtin.RED || nodeColor(node.Right) == builtin.RED {
		return false
	}
	return true
}

func (self *RBTree[T]) BFS2() bool {
	list := []*RBTreeNode[T]{self.Root}
	idx := 0
	end := len(list)
	for idx < end {
		n := list[idx]
		ok := conditionRedNode(n)
		if !ok {
			return false
		}
		if n.Left != nil {
			list = append(list, n.Left)
			end++
		}
		if n.Right != nil {
			list = append(list, n.Right)
			end++
		}
		idx++
	}
	return true
}

// 性质：每个节点的每条路径都包含相同数量的黑色节点
func conditionPath[T comparable](node *RBTreeNode[T]) (int, bool) {
	if nil == node {
		return 0, true
	}
	if nodeColor(node) == builtin.BLACK {
		blackLCount, ok1 := conditionPath(node.Left)
		blackRCount, ok2 := conditionPath(node.Right)
		fmt.Println("黑节点：" + fmt.Sprint(node.V) + " 左边： " + utils.Transform.IntToString(blackLCount) + " 右边： " + utils.Transform.IntToString(blackRCount))
		if blackLCount == blackRCount && ok1 && ok2 {
			return blackLCount + 1, true
		} else {
			return 0, false
		}
	} else {
		blackLCount, ok1 := conditionPath(node.Left)
		blackRCount, ok2 := conditionPath(node.Right)
		fmt.Println("红节点：" + fmt.Sprint(node.V) + " 左边： " + utils.Transform.IntToString(blackLCount) + " 右边： " + utils.Transform.IntToString(blackRCount))
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
		if !ok {
			return false
		}
		if n.Left != nil {
			list = append(list, n.Left)
			end++
		}
		if n.Right != nil {
			list = append(list, n.Right)
			end++
		}
		idx++
	}
	return true
}

func (tree *RBTree[T]) replaceNode(old *RBTreeNode[T], new *RBTreeNode[T]) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

// rotateRight 红黑树树节点右旋，相比AVL树需要处理父节点parent的关系
func (tree *RBTree[T]) rotateRight(root *RBTreeNode[T]) *RBTreeNode[T] {
	left := root.Left
	tree.replaceNode(root, left)
	root.Left = left.Right
	// 如果 y 的右子节点更新其父节点
	if left.Right != nil {
		left.Right.Parent = root
	}
	// 将 root 设置为 y 的右子节点
	left.Right = root
	root.Parent = left
	return left
}

// rotateLeft
func (tree *RBTree[T]) rotateLeft(root *RBTreeNode[T]) *RBTreeNode[T] {
	right := root.Right
	tree.replaceNode(root, right)
	root.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = root
	}
	right.Left = root
	root.Parent = right
	return right
}
