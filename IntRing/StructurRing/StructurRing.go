package StructurRing

import (
	"fmt"
	"sync"
)

//type contElInt struct{
//	//elMutex sync.Mutex
//	el      *int
//}
//
//func (eI *contElInt)write(val *int){
//	//eI.elMutex.Lock()
//	//defer eI.elMutex.Unlock()
//	eI.el = val
//}
//
//func (eI *contElInt)read() *int{
//	//eI.elMutex.Lock()
//	//defer eI.elMutex.Unlock()
//	return eI.el
//}

type IntRing struct{

	conditionSignalFull *sync.Cond
	conditionSignalEmpty *sync.Cond
	indexStartMutex sync.Mutex
	start int
	indexEndMutex sync.Mutex
	end int
	data []*int
	isfull bool
	isempty bool
	size int
	writeChan chan *int
	readChan chan *int
}

func NewIntRing(size, countWriter int) *IntRing {
	if size == 0{
		fmt.Println("Size не может равнятся 0")
		return nil
	}
	ring := &IntRing{

		sync.NewCond(&sync.Mutex{}),
		sync.NewCond(&sync.Mutex{}),
		sync.Mutex{},
		0,
		sync.Mutex{},
		0,
		make([]*int, size,size),
		false,
		true,
		size,
		make(chan *int),
		make(chan *int),
	}
	//for i:= 0 ; i < size; i++{
	//	//ring.data[i] = &contElInt{sync.Mutex{}, nil}
	//	ring.data[i] = &contElInt{}
	//}
	for i:=0; i < countWriter; i++{
		go ring.writer()
	}
	return ring
}

func (r *IntRing) IsFull() bool{
	return r.isfull
}

func (r *IntRing) isEmpty()bool{
	return r.isempty
}

func (r *IntRing) Size() int{
	return r.size
}

func (r * IntRing)writer(){
	for {
		 val := <- r.writeChan
		 r.indexEndMutex.Lock()
		 defer func (){
			 if r.IsFull(){
				 r.conditionSignalFull.L.Lock()
				 r.conditionSignalFull.Wait()
				 r.conditionSignalFull.L.Unlock()
			 }
		 }()

		 fmt.Println("End: ", r.end)
		 r.data[r.end]=val
		 r.end++
		if r.end == r.size {
			r.end = 0
		}
		if r.end == r.start{
			r.isfull = true
		}
		if r.isEmpty(){
			r.isempty = false
			r.conditionSignalEmpty.Broadcast()
		}
		r.indexEndMutex.Unlock()
	}
}

func (r *IntRing)Write(val int) error {
	select {
	case r.writeChan <- &val:
		return nil
	default:
		return fmt.Errorf("Buffer is full")
	}
}

func (r *IntRing)reader(){
		r.indexStartMutex.Lock()
		defer r.indexStartMutex.Unlock()
		if r.isEmpty(){
			r.conditionSignalEmpty.L.Lock()
			r.conditionSignalEmpty.Wait()
			r.conditionSignalEmpty.L.Unlock()
		}
		el := r.data[r.start]
		r.readChan <- el
		fmt.Println("Start: ", r.start)
		if r.start++; r.start == r.size {
			r.start = 0
		}
		if r.start == r.end{
			r.isempty = true
		}
		if r.IsFull(){
			r.isfull = false
			r.conditionSignalFull.Broadcast()
		}
}

func (r *IntRing)Read() int {
	go r.reader()
	return *<-r.readChan
}

