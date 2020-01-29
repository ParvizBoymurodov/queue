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
	if p.firstElement()!=p.first.value{
		t.Error("when no one is in line, no one will be first in line",p.first.value)
	}
	if p.lastElement()!=p.first.value{
		t.Error("when no one is in line no one will be last in line",p.last.value)
	}
}

func Test_addInQueue(t *testing.T) {
	p := queue{}
	p.equeue(1)
	if p.len()!=1{
		t.Error("when there is one person in line then the length of the queue is 1",p.len())
	}
	if p.firstElement() != p.lastElement() {
		t.Error("after one person appears in the queue, the queue length will be 1,  got: ", p.firstElement() ,p.lastElement())
	}
	if p.firstElement()!=p.first.value{
		t.Error("if there is one person in line, then he’s the first",p.firstElement())
	}
	if p.lastElement()!=p.first.value{
		t.Error("if there is one person in line, then he is the first and last",p.lastElement())
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
