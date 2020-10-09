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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOder(t *testing.T) {

	expected := []interface{}{3, 0, 2, 5, 4}

	r := NewBinaryTree(3)
	root := r.Root
	root.Left = AddNode(0)
	root.Left.Right = AddNode(2)
	root.Right = AddNode(5)
	root.Right.Left = AddNode(4)

	actual := r.PreOrder()
	assert.Equal(t, true, r.IsBalanced(), "The two item should be the same.")
	assert.Equal(t, 3, r.High(), "The two item should be the same.")
	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestMiddleOder(t *testing.T) {

	expected := []interface{}{0, 2, 3, 4, 5}

	r := NewBinaryTree(3)
	root := r.Root
	root.Left = AddNode(0)
	root.Left.Right = AddNode(2)
	root.Right = AddNode(5)
	root.Right.Left = AddNode(4)

	actual := r.MiddleOrder()

	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestPostOder(t *testing.T) {

	expected := []interface{}{2, 0, 7, 6, 4, 5, 3}

	r := NewBinaryTree(3)
	root := r.Root
	root.Left = AddNode(0)
	root.Left.Right = AddNode(2)
	root.Right = AddNode(5)
	root.Right.Left = AddNode(4)
	root.Right.Left.Right = AddNode(6)
	root.Right.Left.Right.Right = AddNode(7)

	actual := r.PostOrder()
	assert.Equal(t, false, r.IsBalanced(), "The two item should be the same.")
	assert.Equal(t, 5, r.High(), "The two item should be the same.")
	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestBFS(t *testing.T) {

	expected := []interface{}{3, 0, 5, 2, 4}

	r := NewBinaryTree(3)
	root := r.Root
	root.Left = AddNode(0)
	root.Left.Right = AddNode(2)
	root.Right = AddNode(5)
	root.Right.Left = AddNode(4)

	actual := r.BFS()
	assert.Equal(t, true, r.IsBalanced(), "The two item should be the same.")
	assert.Equal(t, 3, r.High(), "The two item should be the same.")
	assert.Equal(t, expected, actual, "The two item should be the same.")
	assert.Equal(t, 5, r.Size(), "The two item should be the same.")
}
