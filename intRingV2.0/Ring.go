package intRingV2

import "fmt"

type Ring struct {
	size int
	data []*int
	startSyncChan chan int
	start int
	endSyncChan chan int
	end int
	isFull bool
	isEmpty bool
}

func NewRing (size int)*Ring{
	return &Ring{
		size: size,
		data: make([]*int, size, size),
		startSyncChan: make(chan int, 1),
		start: 0,
		endSyncChan: make(chan int, 1),
		end: 0,
		isFull: false,
		isEmpty: true,
	}
}

func (r *Ring)IsFull()bool{
	return r.isFull
}

func (r *Ring)IsEmpty()bool{
	return r.isEmpty
}

func (r *Ring)Write(val int)error{
	r.endSyncChan <- 0
	defer func(){
		<-r.endSyncChan
	}()
	if !r.IsFull(){
		r.data[r.end]=&val
		if r.end++; r.end == r.size{
			r.end = 0
		}
		if r.end ==r.start{
			r.isFull = true
		}
		if r.IsEmpty(){
			r.isEmpty = false
		}
		return nil
	} else {
		return fmt.Errorf("Buffer FULL")
	}
}


func (r *Ring) Read() (int, error){
	r.startSyncChan <- 0
	defer func(){
		<-r.startSyncChan
	}()
	if !r.IsEmpty(){
		val := *r.data[r.start]
		if r.start++; r.start == r.size{
			r.start = 0
		}
		if r.start == r.end{
			r.isEmpty = true
		}
		if r.IsFull(){
			r.isFull = false
		}
		return val, nil
	} else {
		return 0, fmt.Errorf("Buffer EMPTY")
	}
}