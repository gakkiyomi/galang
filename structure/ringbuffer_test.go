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

func TestRingBuffer(t *testing.T) {
	buffer := NewRingBuffer[int](3)
	assert.Equal(t, true, buffer.IsEmpty(), "The ringbuffer should be empty.")
	assert.Equal(t, []int{}, buffer.GetAll(), "The ringbuffer should be empty.")
	buffer.Put(1)
	buffer.Put(2)
	buffer.Put(3)
	assert.Equal(t, true, buffer.IsFull(), "The ringbuffer should be full.")
	assert.Equal(t, 3, buffer.Len(), "The ringbuffer len shoud be 3.")
	assert.Equal(t, 1, buffer.Get(), "The ringbuffer should be 1")
	assert.Equal(t, 2, buffer.Get(), "The two item should be the 2.")
	assert.Equal(t, 3, buffer.Get(), "The two item should be the 3.")
	assert.Equal(t, 0, buffer.Get(), "The two item should be the 0.")
	buffer.Put(4)
	assert.Equal(t, 4, buffer.Get(), "The two item should be the 4.")
	assert.Equal(t, 0, buffer.Len(), "The ringbuffer len shoud be 0")
	buffer.Put(5)
	buffer.Put(6)
	buffer.Put(7)
	assert.Equal(t, false, buffer.Put(8), "The ringbuffer is full cannot append data")
	assert.Equal(t, []int{5, 6, 7}, buffer.GetAll(), "The ringbuffer should be 5,6,7.")
	assert.Equal(t, true, buffer.IsEmpty(), "The ringbuffer should be empty.")
}
