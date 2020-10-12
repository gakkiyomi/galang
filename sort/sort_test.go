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

func TestHeapSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 13, 22, 44, 54, 222, 441}
	expected2 := []int{441, 222, 54, 44, 22, 13, 4, 3, 2, 1}
	source := []int{1, 4, 2, 44, 22, 13, 222, 441, 3, 54}
	source2 := []int{1, 4, 2, 44, 22, 13, 222, 441, 3, 54}
	x := HeapSort(source, true)
	y := HeapSort(source2, false)
	assert.Equal(t, expected, x, "The two item should be the same.")
	assert.Equal(t, expected2, y, "The two item should be the same.")
}

func TestMergeSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 13, 22, 44, 54, 211, 222, 441}
	source := []int{1, 4, 2, 44, 22, 13, 222, 441, 3, 54, 211}

	x := MergeSort(source)
	assert.Equal(t, expected, x, "The two item should be the same.")

}

func TestShellSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 6, 13, 22, 44, 54, 211, 222, 441, 2222, 11114}
	source := []int{1, 4, 2, 44, 22, 13, 222, 441, 3, 54, 11114, 211, 2222, 6}

	ShellSort(source)
	assert.Equal(t, expected, source, "The two item should be the same.")

}
