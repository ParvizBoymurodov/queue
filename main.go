package main

type queue struct {
	first *queueNode
	last  *queueNode
	size  int
}
type queueNode struct {
	next  *queueNode
	prev  *queueNode
	value interface{}
}


func (receiver *queue) len() int {
	return receiver.size
}
func (receiver *queue) firstElement() interface{} {
	return receiver.first.value
}
func (receiver *queue) lastElement() interface{} {
	return receiver.last.value
}
func (receiver *queue) equeue(value interface{}) {
	if receiver.len() == 0 {
		receiver.first = &queueNode{
			next:  nil,
			prev:  nil,
			value: value,
		}
		receiver.last = receiver.first
		receiver.size++
		return
	}
	receiver.size++
	current := receiver.first
	for {
		if current.next == nil {
			current.next = &queueNode{
				next:  nil,
				prev:  current,
				value: value,
			}
			receiver.last = current.next
			return
		}
		current = current.next
	}
}
func (receiver *queue) dequeue() interface{} {
	if receiver.len() == 0 {
		return queue{}
	}
	if receiver.len() == 1 {
		receiver.first = nil
		receiver.last = nil
		receiver.size--
		return queue{}
	}
	receiver.first = receiver.first.next
	receiver.first.prev = nil
	receiver.size--
return queue{}
}
func main() {}
