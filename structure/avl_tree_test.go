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
	"testing"

	"github.com/gakkiyomi/galang/utils"
	"github.com/stretchr/testify/assert"
)

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree(func(v1, v2 int) int {
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
	assert.Equal(t, 0, tree.High(), "The two item should be the same.")
	tree.Insert(99)
	tree.Insert(98)
	tree.Insert(100)
	assert.Equal(t, 98, tree.Root.Left.V, "The two item should be the same.")
	assert.Equal(t, 99, tree.Root.V, "The two item should be the same.")
	assert.Equal(t, 100, tree.Root.Right.V, "The two item should be the same.")
	assert.Equal(t, 1, tree.Root.Left.Height, "The two item should be the same.")
	assert.Equal(t, 1, tree.Root.Right.Height, "The two item should be the same.")
	assert.Equal(t, 2, tree.Root.Height, "The two item should be the same.")
	assert.Equal(t, 100, tree.Max(), "The two item should be the same.")
	assert.Equal(t, 98, tree.Min(), "The two item should be the same.")
	assert.Equal(t, 99, tree.Search(99), "The two item should be the same.")
	assert.Equal(t, 0, tree.Search(1000), "The two item should be the same.")
	tree.Delete(106)
	tree.Delete(100)
	tree.Insert(101)
	tree.Insert(70)
	tree.Insert(102)
	tree.Delete(102)
	tree.Insert(72)
	tree.Insert(103)
	tree.Insert(120)
	tree.Insert(104)
	tree.Insert(105)
	tree.Insert(106)
	fmt.Println(utils.Transform.AnyToString(tree))
	assert.Equal(t, 120, tree.Max(), "The two item should be the same.")
	tree.Delete(106)
	assert.Equal(t, 0, tree.Search(106), "The two item should be the same.")
	tree.Insert(1009)
	tree.Insert(233)
	tree.Insert(1414)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	assert.Equal(t, 1414, tree.Max(), "The two item should be the same.")
	assert.Equal(t, 1, tree.Min(), "The two item should be the same.")
	assert.LessOrEqual(t, GetBalance(tree.Root), 1)
	tree.Delete(72)
	tree.Delete(103)
	tree.Delete(120)
	tree.Delete(104)
	tree.Delete(105)
	tree.Delete(106)
	assert.LessOrEqual(t, GetBalance(tree.Root), 1)
}
