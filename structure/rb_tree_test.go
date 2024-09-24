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

func TestRBTree(t *testing.T) {
	tree := NewRBTree(func(v1, v2 int) int {
		switch {
		case v1 < v2:
			return -1
		case v1 == v2:
			return 0
		default:
			return 1
		}
	})
	assert.Nil(t, tree.Root, "The two item should be the same.")
	tree.Insert(99)
	tree.Insert(99)
	tree.Insert(98)
	tree.Insert(100)
	assert.Equal(t, 98, tree.Root.Left.V, "The two item should be the same.")
	assert.Equal(t, 99, tree.Root.V, "The two item should be the same.")
	assert.Equal(t, 100, tree.Root.Right.V, "The two item should be the same.")
	assert.Equal(t, 100, tree.Max(), "The two item should be the same.")
	assert.Equal(t, 98, tree.Min(), "The two item should be the same.")
	data, exist := tree.Search(99)
	assert.Equal(t, 99, data, "The two item should be the same.")
	assert.True(t, exist, "The two item should be the same.")
	tree.Delete(99)
	data, exist = tree.Search(1000)
	assert.Equal(t, 0, data, "The two item should be the same.")
	assert.False(t, exist, "The two item should be the same.")
	tree.Insert(70)
	tree.Insert(72)
	tree.Insert(101)
	tree.Insert(102)
	tree.Insert(103)
	tree.Insert(120)
	assert.Equal(t, 120, tree.Max(), "The two item should be the same.")
	tree.Insert(104)
	tree.Delete(120)
	_, exist = tree.Search(120)
	assert.False(t, exist, "The two item should be the same.")
	tree.Insert(14)
	tree.Insert(9)
	tree.Insert(55)
	tree.Insert(52)
	tree.Insert(80)
	tree.Insert(105)
	tree.Insert(106)
	data, exist = tree.Search(106)
	assert.Equal(t, 106, data, "The two item should be the same.")
	assert.True(t, exist, "The two item should be the same.")
	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(522)
	tree.Insert(180)
	tree.Delete(5)
	tree.Delete(180)
	tree.Delete(522)
	tree.Delete(105)
	assert.True(t, tree.BFS2(), "The two item should be the same.")
	assert.True(t, tree.BFS(), "The two item should be the same.")
}
