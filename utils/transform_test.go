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

func TestInt(t *testing.T) {
	res1 := Transform.IntToString(5)
	res2, _ := Transform.StringToInt("5")
	assert.Equal(t, "5", res1)
	assert.Equal(t, 5, res2)
}

func TestAnyToString(t *testing.T) {
	assert.Equal(t, "1", Transform.AnyToString(1))
	assert.Equal(t, "1.5", Transform.AnyToString(1.5))
	assert.Equal(t, "7777777777", Transform.Int64ToString(int64(7777777777)))
	assert.Equal(t, "{\"apple\":1,\"banana\":2,\"orange\":3}", Transform.AnyToString(map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}))
}
