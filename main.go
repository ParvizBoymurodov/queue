package main

type queue struct {
	first *listQueue
	last  *listQueue
	size  int
}
type listQueue struct {
	next  *listQueue
	prev  *listQueue
	value interface{}
}

func (Queue *queue) len() int {
	return Queue.size
}
func (Queue *queue) firstElement() *listQueue {
	return Queue.first
}
func (Queue *queue) lastElement() *listQueue {
	return Queue.last
}
func (Queue *queue) equeue(value interface{}) {
	if Queue.len() == 0 {
		Queue.first = &listQueue{
			next:  nil,
			prev:  nil,
			value: value,
		}
		Queue.last = Queue.first
		Queue.size++
		return
	}
	Queue.size++
	current := Queue.first
	for {
		if current.next == nil {
			current.next = &listQueue{
				next:  nil,
				prev:  current,
				value: value,
			}
			Queue.last = current.next
			return
		}
		current = current.next
	}
}
func (Queue *queue) dequeue() {
	if Queue.len() == 0 {
		return
	}
	if Queue.len() == 1 {
		Queue.first = nil
		Queue.last = nil
		Queue.size--
		return
	}
	Queue.first = Queue.first.next
	Queue.first.prev = nil
	Queue.size--

}
func main() {}
