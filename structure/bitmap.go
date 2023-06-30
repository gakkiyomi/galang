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

func NewBitMap() *SingleBitMap {
	return &SingleBitMap{}
}

func (self *SingleBitMap) Has(num int) bool {
	//bitmap 是指num在bitmap数组里是第几个bitmap中: num < 64  bitmap = 0  num >= 64 bitmap =1
	//bit 是指 num在当前bitmap 0 的二进制表示中的index值
	bitmap, bit := num/64, uint64(num%64)
	return bitmap < len(self.bitmaps) && (self.bitmaps[bitmap]&(1<<bit)) != 0
}

func (self *SingleBitMap) Clear() {
	self.bitmaps = nil
	self.len = 0
}

func (self *SingleBitMap) Add(num int) {
	bitmap, bit := num/64, uint64(num%64)
	//如果传进来的数超过bitmap数组的最大值，需要将对bitmap进行扩容
	for bitmap >= len(self.bitmaps) {
		self.bitmaps = append(self.bitmaps, 0)
	}
	// 判断num是否已经存在bitmap中
	if self.bitmaps[bitmap]&(1<<bit) == 0 {
		self.bitmaps[bitmap] |= 1 << bit
		self.len++
	}
}

func (self *SingleBitMap) Len() int {
	return self.len
}

func (self *SingleBitMap) Sort(nums []int) []int {
	for _, item := range nums {
		self.Add(item)
	}
	res := make([]int, 0, len(nums))
	for i, v := range self.bitmaps {
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
