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

	root := NewBinaryTree(3)
	root.Left = NewBinaryTree(0)
	root.Left.Right = NewBinaryTree(2)
	root.Right = NewBinaryTree(5)
	root.Right.Left = NewBinaryTree(4)

	actual := root.PreOrder()
	assert.Equal(t, true, root.IsBalanced(), "The two item should be the same.")
	assert.Equal(t, 3, root.High(), "The two item should be the same.")
	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestMiddleOder(t *testing.T) {

	expected := []interface{}{0, 2, 3, 4, 5}

	root := NewBinaryTree(3)
	root.Left = NewBinaryTree(0)
	root.Left.Right = NewBinaryTree(2)
	root.Right = NewBinaryTree(5)
	root.Right.Left = NewBinaryTree(4)

	actual := root.MiddleOrder()

	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestPostOder(t *testing.T) {

	expected := []interface{}{2, 0, 7, 6, 4, 5, 3}

	root := NewBinaryTree(3)
	root.Left = NewBinaryTree(0)
	root.Left.Right = NewBinaryTree(2)
	root.Right = NewBinaryTree(5)
	root.Right.Left = NewBinaryTree(4)
	root.Right.Left.Right = NewBinaryTree(6)
	root.Right.Left.Right.Right = NewBinaryTree(7)

	actual := root.PostOrder()
	assert.Equal(t, false, root.IsBalanced(), "The two item should be the same.")
	assert.Equal(t, 5, root.High(), "The two item should be the same.")
	assert.Equal(t, expected, actual, "The two item should be the same.")
}
