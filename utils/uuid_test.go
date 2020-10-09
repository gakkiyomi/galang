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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	fmt.Println(UUID.NewUUID())
}

func TestCheckUUID(t *testing.T) {
	b := UUID.Check("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	c := UUID.Check("6ba7b810-9dad-11d1-80b4-00c04fd430c8aaa")
	assert.Equal(t, true, b, "The two item should be the same.")
	assert.Equal(t, false, c, "The two item should be the same.")
}
