package sq

import "sync"

// Queue represents a thread-safe FIFO queue
type Queue struct {
	mu    sync.Mutex
	items []Item
}

// Enqueue adds an item to the end of the queue (RPUSH)
func (q *Queue) Enqueue(item Item) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Dequeue removes an item from the front of the queue (LPOP)
func (q *Queue) Dequeue() (Item, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Length returns the number of items in the queue (LLEN)
func (q *Queue) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

// EnqueueFront adds an item to the front of the queue (LPUSH)
func (q *Queue) EnqueueFront(item Item) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append([]Item{item}, q.items...)
}

// DequeueBack removes an item from the end of the queue (RPOP)
func (q *Queue) DequeueBack() (Item, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return item, true
}
