package queue

import "fmt"

type Element struct {
	e interface{}
	next *Element
}

type Queue struct {
	head *Element
	tail *Element
}

func New() *Queue {
	q := new(Queue)
	q.head = nil
	q.tail = nil
	return q
}

func (Q *Queue) Poll() interface{} {
	if Q.head != nil {
		t := Q.head
		Q.head = Q.head.next
		t.next = nil
		return t.e
	}
	return nil
}

func (Q *Queue) Push(v interface{}) {
	e := Element {
		e : v,
		next: nil,
	}
	if Q.tail == nil {
		Q.tail = &e
	} else {
		Q.tail.next = &e
		Q.tail = Q.tail.next
	}
	if Q.head == nil {
		Q.head = Q.tail
	}
}

func (Q *Queue) DebugInfo() {
	var cnt int = 0
	var t *Element = Q.head
	for {
		if t == nil {
			break
		}
		cnt = cnt + 1
		t = t.next
	}
	fmt.Println("queue num:", cnt)
}