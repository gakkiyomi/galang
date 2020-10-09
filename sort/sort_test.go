// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 13, 22, 44, 54, 222, 441}
	source := []int{1, 4, 2, 44, 22, 13, 222, 441, 3, 54}
	BubbleSort(source)
	assert.Equal(t, expected, source, "The two item should be the same.")
}

func TestSelectionSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 13, 22, 44, 54, 222, 441}
	source := []int{1, 4, 2, 44, 22, 13, 222, 441, 3, 54}
	SelectionSort(source)
	assert.Equal(t, expected, source, "The two item should be the same.")
}

func TestQuickSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 13, 22, 44, 54, 222, 441}
	source := []int{1, 4, 2, 44, 22, 13, 222, 441, 3, 54}
	QuickSort(source)
	assert.Equal(t, expected, source, "The two item should be the same.")
}
