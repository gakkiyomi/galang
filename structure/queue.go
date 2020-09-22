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

type Queue struct {
	head *node
	tail *node
	len  int
}

func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

func (q *Queue) Offer(v interface{}) {
	n := &node{nil, v, nil}
	if q.len == 0 {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = q.tail.next
	}
	q.len++
}

func (q *Queue) Poll() interface{} {
	if q.len == 0 {
		return nil
	}
	res := q.head
	q.head = q.head.next
	q.len--
	return res.v

}

//retrun head node value ,if queue isEmpty return nil
func (q *Queue) Peek() interface{} {
	if q.len == 0 {
		return nil
	}
	return q.head.v
}

func (q *Queue) Len() int {
	return q.len
}
