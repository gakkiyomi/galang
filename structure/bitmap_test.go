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

func TestBitMap(t *testing.T) {
	bitmap := NewBitMap()
	assert.Equal(t, 0, bitmap.Len(), "The bitmap should be empty.")
	assert.Equal(t, false, bitmap.Has(5), "The bitmap should be empty.")
	bitmap.Add(3)
	bitmap.Add(2)
	bitmap.Add(444)
	assert.Equal(t, true, bitmap.Has(444), "The reslut should be true.")
	assert.Equal(t, false, bitmap.Has(44), "The reslut should be true.")
	assert.Equal(t, true, bitmap.Has(2), "The reslut should be true.")
	assert.Equal(t, true, bitmap.Has(3), "The reslut should be true.")
	assert.Equal(t, 3, bitmap.Len(), "The bitmap len should be 3.")
	assert.Equal(t, 7, len(bitmap.bitmaps), "The bitmap len should be 3.")
	bitmap.Clear()
	assert.Equal(t, 0, bitmap.Len(), "The bitmap should be empty.")
	assert.Equal(t, []int{1, 2, 3, 4, 55, 3333}, bitmap.Sort([]int{3, 4, 1, 2, 3333, 55}), "The reslut should be true.")
}
