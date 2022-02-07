package Graf

import "fmt"

type Que struct {
	first *El
	end *El
	size int
}

type El struct{
	element interface{}
	next *El
}

func (q *Que)Push(el interface{}){
	if q.first == nil {
		q.first = &El{el, nil}
		q.end = q.first
	}else{
		(*q.end).next = &El{el, nil}
		q.end = (*q.end).next
	}
	q.size++
}

func (q *Que)Pop()(interface{}, error){
	if q.first == nil {
		return nil, fmt.Errorf("Queue empty")
	}
	var retEl *El
	retEl, q.first = q.first, q.first.next
	q.size--
	return retEl.element, nil
}

func NewQue() *Que{
	return &Que{nil,nil,0}
}
