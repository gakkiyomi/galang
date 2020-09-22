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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Offer(1)
	queue.Offer(2)
	queue.Offer(nil)
	queue.Offer(3)
	assert.Equal(t, 1, queue.Poll(), "The two item should be the same.")
	assert.Equal(t, 2, queue.Poll(), "The two item should be the same.")
	assert.Equal(t, nil, queue.Poll(), "The two item should be the same.")
	assert.Equal(t, 3, queue.Poll(), "The two item should be the same.")
	assert.Equal(t, nil, queue.Peek(), "The two item should be the same.")
}
