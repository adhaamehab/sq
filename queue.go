package sq

import "sync"

// Queue represents a thread-safe FIFO queue
type Queue struct {
	mu    sync.Mutex
	items []Item
}

// Push adds an item to the end of the queue
func (q *Queue) Push(item Item) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Pop removes an item from the front of the queue
func (q *Queue) Pop() (Item, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Length returns the number of items in the queue
func (q *Queue) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

// PushFront adds an item to the front of the queue
func (q *Queue) PushFront(item Item) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append([]Item{item}, q.items...)
}

// PopBack removes an item from the end of the queue
func (q *Queue) PopBack() (Item, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return item, true
}
