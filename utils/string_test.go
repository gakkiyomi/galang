// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBlank(t *testing.T) {

	a := String.IsBlank("")
	b := String.IsBlank(" ")
	c := String.IsBlank("     ")
	d := String.IsBlank("  aaa   ")
	e := String.IsBlank("     aaa")
	f := String.IsBlank("aaa")

	assert.Equal(t, true, a, "The two item should be the same.")
	assert.Equal(t, true, b, "The two item should be the same.")
	assert.Equal(t, true, c, "The two item should be the same.")
	assert.Equal(t, false, d, "The two item should be the same.")
	assert.Equal(t, false, e, "The two item should be the same.")
	assert.Equal(t, false, f, "The two item should be the same.")

}

func TestToStringArray(t *testing.T) {

	a := String.ToStringArray("")
	b := String.ToStringArray(" ")
	c := String.ToStringArray("     ")
	d := String.ToStringArray("  aaa   ")
	e := String.ToStringArray("     aaa")
	f := String.ToStringArray("aaa")

	assert.Equal(t, 0, len(a), "The two item should be the same.")
	assert.Equal(t, 1, len(b), "The two item should be the same.")
	assert.Equal(t, 5, len(c), "The two item should be the same.")
	assert.Equal(t, 8, len(d), "The two item should be the same.")
	assert.Equal(t, 8, len(e), "The two item should be the same.")
	assert.Equal(t, 3, len(f), "The two item should be the same.")

}

func TestRemoveDuplicateInArray(t *testing.T) {
	expected := []string{`1`, `2`, `a`, `abvvv`}

	source := []string{`1`, `2`, `a`, `a`, `2`, `abvvv`, `a`, `abvvv`}

	actual := String.RemoveDuplicateInArray(source)

	for i := 0; i < 1000000; i++ {
		assert.Equal(t, expected, actual, "The two item should be the same.")
	}

}

func TestStringInsert(t *testing.T) {
	expected := []string{`1`, `2`, `5`, `3`, `4`}

	source := []string{`1`, `2`, `3`, `4`}

	actual := String.InsertAtIndex(source, `5`, 2)

	assert.Equal(t, expected, actual, "The two item should be the same.")
}
