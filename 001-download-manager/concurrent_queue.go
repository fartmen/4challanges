package main

import (
	"sync"
)

type Q struct {
	sync.Mutex
	List []interface{}
}

func (q *Q) Push(s interface{}) {
	q.Lock()
	defer q.Unlock()

	q.List = append(q.List, s)
}

func (q *Q) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.List) == 0 {
		return ""
	}
	item := q.List[0]
	q.List = q.List[1:]
	return item
}

// returns a list with all items currently
// in queue (snapshot)
func (q *Q) Snap() []interface{} {
	res := []interface{}{}

	q.Lock()
	defer q.Unlock()

	for i := 0; i < len(q.List); i++ {
		res = append(res, q.List[i])
	}
	return res
}

func NewQ(capacity int) *Q {
	return &Q{
		List: make([]interface{}, 0, capacity),
	}
}
