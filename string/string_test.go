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

func TestStringBuilder(t *testing.T) {

	builder := String.NewStringBuilder("fangcong")
	builder.Append("\r\n").Append("哈哈哈")
	s := builder.ToString()
	expected := "fangcong\r\n" + "哈哈哈"
	assert.Equal(t, expected, s, "The two item should be the same.")
}
