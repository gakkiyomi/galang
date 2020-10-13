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
