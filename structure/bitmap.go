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

func NewBitMap() *BitMap {
	return &BitMap{}
}

func (bm *BitMap) Has(num int) bool {
	//bitmap 是指num在bitmap数组里是第几个bitmap中: num < 64  bitmap = 0  num >= 64 bitmap =1
	//bit 是指 num在当前bitmap 0 的二进制表示中的index值
	bitmap, bit := num/64, uint64(num%64)
	return bitmap < len(bm.bitmaps) && (bm.bitmaps[bitmap]&(1<<bit)) != 0
}

func (bitmap *BitMap) Clear() {
	bitmap.bitmaps = nil
	bitmap.len = 0
}

func (bm *BitMap) Add(num int) {
	bitmap, bit := num/64, uint64(num%64)
	//如果传进来的数超过bitmap数组的最大值，需要将对bitmap进行扩容
	len := len(bm.bitmaps)
	if bitmap >= len {
		offset := bitmap - len + 1
		bm.bitmaps = append(bm.bitmaps, make([]uint64, offset)...)
	}
	// 判断num是否已经存在bitmap中
	if bm.bitmaps[bitmap]&(1<<bit) == 0 {
		//如果不存在，通过位或运算加入到bitmap中
		bm.bitmaps[bitmap] |= 1 << bit
		bm.len++
	}
}

func (bitmap *BitMap) Len() int {
	return bitmap.len
}

// Sort a array use bitmap
func (bitmap *BitMap) Sort(nums []int) []int {
	for _, item := range nums {
		bitmap.Add(item)
	}
	res := make([]int, 0, len(nums))
	for i, v := range bitmap.bitmaps {
		if v == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if v&(1<<j) != 0 {
				res = append(res, 64*i+j)
			}
		}
	}
	return res
}
