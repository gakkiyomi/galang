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

// NewRingBuffer return a instance of RingBuffer
func NewRingBuffer[T comparable](cap int) RingBuffer[T] {
	buffer := make([]T, cap)
	return &ArrayRingBuffer[T]{buffer, cap, 0, 0, 0}
}

// Len the data size in RingBuffer
func (buffer *ArrayRingBuffer[T]) Len() int {
	return buffer.len
}

// IsEmpty the buffer is empty
func (buffer *ArrayRingBuffer[T]) IsEmpty() bool {
	return (buffer.len == 0)
}

// IsFull the buffer is full
func (buffer *ArrayRingBuffer[T]) IsFull() (isFull bool) {
	isFull = (buffer.len == buffer.cap)
	return
}

// Put put a data to RingBuffer
func (buffer *ArrayRingBuffer[T]) Put(data T) (isSuccess bool) {
	if buffer.IsFull() {
		isSuccess = false
	} else {
		buffer.buffer[buffer.nextIndex(buffer.wirteIndex)] = data
		buffer.wirteIndex++
		buffer.len++
		isSuccess = true
	}
	return
}

// Get get oldest data from RingBuffer
func (buffer *ArrayRingBuffer[T]) Get() (data T) {
	if buffer.IsEmpty() {
		return data
	}
	data = buffer.buffer[buffer.nextIndex(buffer.readIndex)]
	buffer.len--
	buffer.readIndex++
	return data
}

// GetAll get all data from RingBuffer
func (buffer *ArrayRingBuffer[T]) GetAll() (data []T) {
	if buffer.IsEmpty() {
		data = make([]T, 0)
	}
	for !buffer.IsEmpty() {
		data = append(data, buffer.Get())
	}
	return
}

// nextIndex cacl the write/read pointer next index
func (buffer *ArrayRingBuffer[T]) nextIndex(index int) int {
	return (index + 1) % buffer.cap
}
