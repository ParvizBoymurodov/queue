package main

import (
	"testing"
)

func Test_emptyQueue(t *testing.T) {
	p := queue{}
	p.equeue(0)
	if p.len() != 1 {
		t.Error("queue length must be 0, got: ", p.len())
	}
}

func Test_addInQueue(t *testing.T) {
	p := queue{}
	p.equeue(1)
	if p.firstElement() != p.lastElement() {
		t.Error("after one person appears in the queue, the queue length will be 1,  got: ", p.len())
	}
	p.dequeue()
	if p.len() != 0 {
		t.Error("want 0, got:", p.len())
	}
}

func Test_addSeveralQueue(t *testing.T) {

	p := queue{}
	p.equeue(1)
	p.equeue(2)
	p.equeue(3)

	if p.len() != 3 {
		t.Error("after adding three elements, the length of the queue should be 3, got:  ", p.len())
	}

	p.dequeue()
	if p.len() != 2 {
		t.Error("after deleting one element, the queue length should be 2,  got:  ", p.len())
	}
}
