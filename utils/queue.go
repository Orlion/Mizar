package utils

import (
	"container/list"
	"errors"
	"sync"
)

type Queue struct {
	list  *list.List
	mutex sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (queue *Queue) Push(data interface{}) {
	if data == nil {
		return
	}
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.list.PushBack(data)
}

func (queue *Queue) PushQueue(q *Queue) {
	if q == nil {
		return
	}
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.list.PushBackList(q.list)
}

func (queue *Queue) Pop() (interface{}, error) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	if element := queue.list.Front(); element != nil {
		queue.list.Remove(element)
		return element.Value, nil
	}
	return nil, errors.New("pop failed")
}

func (queue *Queue) Clear() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	for element := queue.list.Front(); element != nil; {
		elementNext := element.Next()
		queue.list.Remove(element)
		element = elementNext
	}
}

func (queue *Queue) Len() int {
	return queue.list.Len()
}
