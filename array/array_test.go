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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateInArray(t *testing.T) {
	expected := []string{`1`, `2`, `a`, `abvvv`}

	source := []string{`1`, `2`, `a`, `a`, `2`, `abvvv`, `a`, `abvvv`}

	actual := Array.RemoveDuplicateInStringArray(source)

	for i := 0; i < 1000000; i++ {
		assert.Equal(t, expected, actual, "The two item should be the same.")
	}

}

func TestStringInsert(t *testing.T) {
	expected := []string{`1`, `2`, `5`, `3`, `4`}

	source := []string{`1`, `2`, `3`, `4`}

	actual := Array.InsertAtIndexByStringArray(source, `5`, 2)

	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestIntInsert(t *testing.T) {
	expected := []int{1, 2, 5, 3, 4}

	source := []int{1, 2, 3, 4}

	actual := Array.InsertAtIndexByIntArray(source, 5, 2)

	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestToStringArray(t *testing.T) {

	a := Array.ToStringArray("")
	b := Array.ToStringArray(" ")
	c := Array.ToStringArray("     ")
	d := Array.ToStringArray("  aaa   ")
	e := Array.ToStringArray("     aaa")
	f := Array.ToStringArray("aaa")

	assert.Equal(t, 0, len(a), "The two item should be the same.")
	assert.Equal(t, 1, len(b), "The two item should be the same.")
	assert.Equal(t, 5, len(c), "The two item should be the same.")
	assert.Equal(t, 8, len(d), "The two item should be the same.")
	assert.Equal(t, 8, len(e), "The two item should be the same.")
	assert.Equal(t, 3, len(f), "The two item should be the same.")

}

func TestReverse(t *testing.T) {
	source := []string{`a`, `b`, `c`, `d`}
	expected := []string{`d`, `c`, `b`, `a`}
	Array.Reverse(source)
	assert.Equal(t, expected, source, "The two item should be the same.")
}

func TestGetMaxInArray(t *testing.T) {
	source := []int{33, 2, 1112, 44122}
	expected := 44122
	actual := Array.GetMaxInArray(source)
	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestGetMinInArray(t *testing.T) {
	source := []int{33, 2, 1112, 44122}
	expected := 2
	actual := Array.GetMinInArray(source)
	assert.Equal(t, expected, actual, "The two item should be the same.")
}

func TestGetMinAndMax(t *testing.T) {
	source := []int{33, 2, 1112, 44122}
	expectedMin := 2
	expectedMax := 44122
	min, max := Array.GetMinAndMax(source)
	assert.Equal(t, expectedMin, min, "The two item should be the same.")
	assert.Equal(t, expectedMax, max, "The two item should be the same.")
}

func TestGetUnionForInt(t *testing.T) {
	source := []int{33, 2, 1112, 44122}
	source2 := []int{22, 33, 4, 5, 6}
	expected := []int{33, 2, 1112, 44122, 22, 4, 5, 6}

	res := Array.GetUnionForInt(source, source2)
	assert.Equal(t, expected, res, "The two item should be the same.")

}

func TestGetUnionForString(t *testing.T) {
	source := []string{`33`, `2`, `1112`, `44122`}
	source2 := []string{`22`, `33`, `4`, `5`, `6`}
	expected := []string{`33`, `2`, `1112`, `44122`, `22`, `4`, `5`, `6`}

	res := Array.GetUnionForString(source, source2)
	assert.Equal(t, expected, res, "The two item should be the same.")
}

func TestGetIntersectForInt(t *testing.T) {
	source := []int{33, 2, 1112, 44122}
	source2 := []int{22, 33, 1112, 44122, 6}
	expected := []int{33, 1112, 44122}

	res := Array.GetIntersectForInt(source, source2)
	assert.Equal(t, expected, res, "The two item should be the same.")

}

func TestGetIntersectForString(t *testing.T) {
	source := []string{`33`, `2`, `1112`, `44122`}
	source2 := []string{`22`, `33`, `4`, `5`, `6`}
	expected := []string{`33`}

	res := Array.GetIntersectForString(source, source2)
	assert.Equal(t, expected, res, "The two item should be the same.")
}

func TestGetComplementForString(t *testing.T) {
	source := []string{`1`, `2`, `3`, `4`}
	source2 := []string{`3`, `4`, `5`, `6`}
	res := Array.GetComplementForString(source, source2)
	assert.Equal(t, 4, len(res))

	source3 := []string{`1`, `2`, `3`, `4`}
	source4 := []string{`5`, `7`, `8`, `51`}
	res2 := Array.GetComplementForString(source3, source4)
	assert.Equal(t, 8, len(res2))

	source5 := []string{`1`, `2`, `3`, `4`}
	source6 := []string{`1`, `2`, `3`, `4`}
	res3 := Array.GetComplementForString(source5, source6)
	assert.Equal(t, 0, len(res3))

	source7 := []string{`1`, `2`, `3`, `4`}
	source8 := []string{`1`, `2`, `3`}
	res4 := Array.GetComplementForString(source7, source8)
	assert.Equal(t, 1, len(res4))
}

func TestGetComplementForInt(t *testing.T) {
	source := []int{1, 2, 3, 4}
	source2 := []int{3, 4, 5, 6}
	res := Array.GetComplementForInt(source, source2)
	assert.Equal(t, 4, len(res))

	source3 := []int{1, 2, 3, 4}
	source4 := []int{5, 7, 8, 51}
	res2 := Array.GetComplementForInt(source3, source4)
	assert.Equal(t, 8, len(res2))

	source5 := []int{1, 2, 3, 4}
	source6 := []int{1, 2, 3, 4}
	res3 := Array.GetComplementForInt(source5, source6)
	assert.Equal(t, 0, len(res3))

	source7 := []int{1, 2, 3, 4}
	source8 := []int{1, 2, 3}
	res4 := Array.GetComplementForInt(source7, source8)
	assert.Equal(t, 1, len(res4))
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 4, 6, 11, 22, 44, 52, 61, 77, 88, 222, 444, 555, 6777, 111414}
	bool := Array.BinraySearch(arr, 444)
	assert.Equal(t, true, bool)
}
