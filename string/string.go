// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package string

import (
	"strings"
)

type (
	GalangString byte
)

var (
	String GalangString
)

//StartWith returns string s isStartWith suffix
func (*GalangString) StartWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

//EndWith returns string s isEndWith suffix
func (*GalangString) EndWith(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

//ContainsIgnoreCase cotains ingore case
func (*GalangString) ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

//ToStringArray Converts this string to a new string array.
func (*GalangString) ToStringArray(str string) []string {

	src := []rune(str)
	array := make([]string, len(src))
	for k, v := range src {
		array[k] = string(v)
	}
	return array
}

//IsBlank check a string isblank
func (*GalangString) IsBlank(str string) bool {

	if len(str) == 0 {
		return true
	}
	src := []rune(str)
	for _, v := range src {
		x := string(v)
		if x != " " {
			return false
		}
	}
	return true
}

//IsNotBlank check a string is not blank
func (gs *GalangString) IsNotBlank(str string) bool {
	return !gs.IsBlank(str)
}

//RemoveDuplicateInArray Remove duplicate strings from the given array.
func (*GalangString) RemoveDuplicateInArray(source []string) []string {
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
func (*GalangString) InsertAtIndex(src []string, v string, index int) (res []string) {
	res = append(src, "")
	copy(res[index+1:], res[index:])
	res[index] = v
	return
}
