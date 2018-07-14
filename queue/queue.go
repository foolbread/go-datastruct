package queue

import (
	"errors"
)

var errOutRange error = errors.New("Out of range!")
var errNullHead error = errors.New("head is nil!")

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	r := new(Queue)

	return r
}

func (q *Queue) QueueEnqueue(d int) {
	q.data = append(q.data, d)
}

func (q *Queue) QueueDequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, errOutRange
	}

	r := q.data[0]
	q.data = q.data[1:]
	return r, nil
}

func (q *Queue) QueueHead() (int, error) {
	if len(q.data) == 0 {
		return 0, errNullHead
	}

	return q.data[0], nil
}

func (q *Queue) QueueSize() int {
	return len(q.data)
}
