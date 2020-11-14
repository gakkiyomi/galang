// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package array

type GalangArray byte

var Array GalangArray

//RemoveDuplicateInArray Remove duplicate strings from the given array.
func (*GalangArray) RemoveDuplicateInStringArray(source []string) []string {
	var res []string

	if len(source) == 0 {
		return source
	}

	hash := make(map[string]int, len(source))
	for k, v := range source {
		if _, ok := hash[v]; ok {

		} else {
			res = append(res, v)
		}
		hash[v] = k
	}
	return res
}

//InsertAtIndex Insert a value in a string slice at a given index
func (*GalangArray) InsertAtIndexByStringArray(src []string, v string, index int) (res []string) {
	res = append(src, "")
	copy(res[index+1:], res[index:])
	res[index] = v
	return
}

//InsertAtIndex Insert a value in a int slice at a given index
func (*GalangArray) InsertAtIndexByIntArray(src []int, v int, index int) (res []int) {
	res = append(src, -1)
	copy(res[index+1:], res[index:])
	res[index] = v
	return
}

//ToStringArray Converts this string to a new string array.
func (*GalangArray) ToStringArray(str string) []string {

	src := []rune(str)
	array := make([]string, len(src))
	for k, v := range src {
		array[k] = string(v)
	}
	return array
}

//Reverse array
func (*GalangArray) Reverse(source []string) {
	len := len(source)
	if len == 0 {
		return
	}
	for i := 0; i < len/2; i++ {
		source[i], source[len-i-1] = source[len-i-1], source[i]
	}
}

//GetMaxInArray 获取数组最大值
func (*GalangArray) GetMaxInArray(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//GetMinInArray 获取数组最小值
func (*GalangArray) GetMinInArray(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func (*GalangArray) GetMinAndMax(arr []int) (int, int) {
	len := len(arr)
	min := arr[0]
	max := arr[len-1]
	for i := 1; i < len; i++ {
		if arr[i] < min {
			min = arr[i]
		}

		if arr[len-1-i] > max {
			max = arr[len-1-i]
		}
	}
	return min, max
}

//GetUnionForInt  获取两个int数组之间的并集
func (*GalangArray) GetUnionForInt(arr1, arr2 []int) []int {
	m := make(map[int]struct{}, len(arr1))
	for _, v := range arr1 {
		m[v] = struct{}{}
	}
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			continue
		}
		arr1 = append(arr1, v)
	}
	return arr1
}

//GetUnionForString   获取两个字符串数组之间的并集
func (*GalangArray) GetUnionForString(arr1, arr2 []string) []string {
	m := make(map[string]struct{}, len(arr1))
	for _, v := range arr1 {
		m[v] = struct{}{}
	}
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			continue
		}
		arr1 = append(arr1, v)
	}
	return arr1
}

//GetIntersectForInt  获取两个int数组之间的交集
func (*GalangArray) GetIntersectForInt(arr1, arr2 []int) []int {

	res := make([]int, 0)

	m := make(map[int]struct{}, len(arr1))
	for _, v := range arr1 {
		m[v] = struct{}{}
	}
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}

	}
	return res
}

//GetIntersectForString   获取两个字符串数组之间的交集
func (*GalangArray) GetIntersectForString(arr1, arr2 []string) []string {
	res := make([]string, 0)

	m := make(map[string]struct{}, len(arr1))
	for _, v := range arr1 {
		m[v] = struct{}{}
	}
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}

	}
	return res
}
