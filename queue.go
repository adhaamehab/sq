package sq

import "sync"

// Queue represents a thread-safe FIFO queue
type Queue struct {
	mu    sync.Mutex
	items []Item
}

// Append adds an item to the end of the queue
func (q *Queue) Append(item Item) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Length returns the number of items in the queue
func (q *Queue) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}
