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
	"github.com/songtianyi/rrframework/logs"
)

func (hp *Heap) GetLeftIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (hp *Heap) GetRightIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (hp *Heap) GetParentIndex(index int) int {
	return (index - 1) / 2
}

func (hp *Heap) GetLeft(parentIndex int) int {
	return hp.Items[hp.GetLeftIndex(parentIndex)]
}

func (hp *Heap) GetRight(parentIndex int) int {
	return hp.Items[hp.GetRightIndex(parentIndex)]
}

func (hp *Heap) GetParent(index int) int {
	return hp.Items[hp.GetParentIndex(index)]
}

func (hp *Heap) HasLeft(index int) bool {
	return hp.GetLeftIndex(index) < len(hp.Items)
}

func (hp *Heap) HasRight(index int) bool {
	return hp.GetRightIndex(index) < len(hp.Items)
}

func (hp *Heap) HasParent(index int) bool {
	return hp.GetParentIndex(index) >= 0
}

func (hp *Heap) Swap(index1, index2 int) {
	hp.Items[index1], hp.Items[index2] = hp.Items[index2], hp.Items[index1]
}

// -------------------------max heap----------------------------

type MaxHeap struct {
	*Heap
}

func NewMaxHeap(source []int) *MaxHeap {
	h := &MaxHeap{
		&Heap{
			Items: source,
		},
	}

	if len(h.Items) > 0 {
		h.buildMaxHeap()
	}

	return h
}

//对于叶子节点，不用调整次序，根据满二叉树的性质，叶子节点比内部节点的个数多1.所以i=n/2 -1 ，不用从n开始。
func (h *MaxHeap) buildMaxHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *MaxHeap) Insert(item int) *MaxHeap {
	h.Items = append(h.Items, item)
	h.shiftUp(len(h.Items) - 1)
	return h
}

func (h *MaxHeap) ExtractMax() int {
	if len(h.Items) == 0 {
		logs.Error("No items in the heap")
	}
	minItem := h.Items[0]

	h.Items[0] = h.Items[len(h.Items)-1]

	h.Items = h.Items[:len(h.Items)-1]

	h.shiftDown(0)

	return minItem
}

func (h *MaxHeap) shiftUp(index int) {
	for h.HasParent(index) && h.GetParent(index) < h.Items[index] {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}

func (h *MaxHeap) shiftDown(index int) {

	//如果当前index存在左或者右节点并且大于当前节点的值则需要交换
	for (h.HasLeft(index) && h.Items[index] < h.GetLeft(index)) || (h.HasRight(index) && h.Items[index] < h.GetRight(index)) {
		//如果左右节点都大于父节点
		if (h.HasLeft(index) && h.Items[index] < h.GetLeft(index)) && (h.HasRight(index) && h.Items[index] < h.GetRight(index)) {
			//找到较大的一个的进行交换
			if h.GetLeft(index) > h.GetRight(index) {
				h.Swap(index, h.GetLeftIndex(index))
				index = h.GetLeftIndex(index)
			} else {
				h.Swap(index, h.GetRightIndex(index))
				index = h.GetRightIndex(index)
			}
		} else if h.HasLeft(index) && h.Items[index] < h.GetLeft(index) { //只有左节点大于父节点
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
		} else { //只有右节点大于父节点
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}

// -------------------------max heap----------------------------

// -------------------------min heap----------------------------

type MinHeap struct {
	*Heap
}

func NewMinHeap(source []int) *MinHeap {
	h := &MinHeap{
		&Heap{
			Items: source,
		},
	}

	if len(h.Items) > 0 {
		h.buildMinHeap()
	}

	return h
}

//对于叶子节点，不用调整次序，根据满二叉树的性质，叶子节点比内部节点的个数多1.所以i=n/2 -1 ，不用从n开始。
func (h *MinHeap) buildMinHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *MinHeap) Insert(item int) *MinHeap {
	h.Items = append(h.Items, item)
	h.shiftUp(len(h.Items) - 1)
	return h
}

func (h *MinHeap) ExtractMin() int {
	if len(h.Items) == 0 {
		logs.Error("No items in the heap")
	}
	minItem := h.Items[0]

	h.Items[0] = h.Items[len(h.Items)-1]

	h.Items = h.Items[:len(h.Items)-1]

	h.shiftDown(0)

	return minItem
}

func (h *MinHeap) shiftUp(index int) {
	for h.HasParent(index) && h.GetParent(index) > h.Items[index] {
		h.Swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
	}
}

func (h *MinHeap) shiftDown(index int) {

	//如果当前index存在左或者右节点并且小于当前节点的值则需要交换
	for (h.HasLeft(index) && h.Items[index] > h.GetLeft(index)) || (h.HasRight(index) && h.Items[index] > h.GetRight(index)) {
		//如果左右节点都小于父节点
		if (h.HasLeft(index) && h.Items[index] > h.GetLeft(index)) && (h.HasRight(index) && h.Items[index] > h.GetRight(index)) {
			//找到较小的一个的进行交换
			if h.GetLeft(index) < h.GetRight(index) {
				h.Swap(index, h.GetLeftIndex(index))
				index = h.GetLeftIndex(index)
			} else {
				h.Swap(index, h.GetRightIndex(index))
				index = h.GetRightIndex(index)
			}
		} else if h.HasLeft(index) && h.Items[index] > h.GetLeft(index) { //只有左节点小于父节点
			h.Swap(index, h.GetLeftIndex(index))
			index = h.GetLeftIndex(index)
		} else { //只有右节点小于父节点
			h.Swap(index, h.GetRightIndex(index))
			index = h.GetRightIndex(index)
		}
	}
}

// -------------------------min heap----------------------------
