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

//NewQueue retruns a Queue
func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

//Offer a value into queue
func (q *Queue) Offer(v interface{}) {
	n := &node{nil, v, nil}
	if q.len == 0 {
		q.tail = n
		q.head = n
	} else {
		n.pre = q.tail
		q.tail.next = n
		q.tail = q.tail.next
	}
	q.len++
}

//Poll a value out of queue
func (q *Queue) Poll() interface{} {
	if q.len == 0 {
		return nil
	}
	res := q.head
	q.head = q.head.next
	q.len--
	return res.v

}

//Peek retrun head node value ,if queue isEmpty return nil
func (q *Queue) Peek() interface{} {
	if q.len == 0 {
		return nil
	}
	return q.head.v
}

//Len returns the queue length
func (q *Queue) Len() int {
	return q.len
}

//IsEmpty return this queue isempty
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

//------------------------Deque------------------------------

//OfferLeft  insert a value into queue head
func (q *Queue) OfferLeft(v interface{}) {
	n := &node{nil, v, nil}
	if q.len == 0 {
		q.tail = n
		q.head = n
	} else {
		n.next = q.head
		q.head.pre = n
		q.head = q.head.pre
	}
	q.len++
}

//PollLeft  delete a value into queue head
func (q *Queue) PollLeft() interface{} {
	if q.len == 0 {
		return nil
	}
	res := q.head
	q.head = q.head.next
	q.len--
	return res.v
}

//Left get head in queue
func (q *Queue) Left() interface{} {
	if q.len == 0 {
		return nil
	}
	return q.head.v
}

//OfferRight  insert a value into queue tail
func (q *Queue) OfferRight(v interface{}) {
	n := &node{nil, v, nil}
	if q.len == 0 {
		q.tail = n
		q.head = n
	} else {
		n.pre = q.tail
		q.tail.next = n
		q.tail = q.tail.next
	}
	q.len++
}

//PollRight  delete a value into queue tail
func (q *Queue) PollRight() interface{} {
	if q.len == 0 {
		return nil
	}
	res := q.tail
	q.tail = q.tail.pre
	q.len--
	return res.v
}

//Right get last item in queue
func (q *Queue) Right() interface{} {
	if q.len == 0 {
		return nil
	}
	return q.tail.v
}
