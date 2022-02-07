package structurDMX

import (
	"fmt"
	"math"
	"sync"
)

type Message string

type DMX struct {
	limit       float64
	inDmxChan   chan Message
	MtRwDmxChan sync.RWMutex
	DmxChans    []chan<- Message

}

func NewDMXUnlimit() *DMX {
	return &DMX{math.Inf(1),
		make(chan Message),
		sync.RWMutex{},
		make([]chan <- Message, 0),
	}
}

func NewDMXLimit(limit int) *DMX {
	return &DMX{float64(limit),
		make(chan Message),
		sync.RWMutex{},
		make([]chan <- Message, 0),
	}
}

func (d *DMX) AddChan(ch... chan<- Message) error{
	if float64(len(d.DmxChans)+len(ch)) > d.limit{
		return fmt.Errorf("DMX is full")
	}

	d.MtRwDmxChan.Lock()
	defer d.MtRwDmxChan.Unlock()
	d.DmxChans = append(d.DmxChans, ch...)

	return nil
}

//func (d *DMX) AddChan(ch chan<- Message) error{
//	if float64(len(d.DmxChans)+1) > d.limit{
//		return fmt.Errorf("DMX is full")
//	}
//
//	d.MtRwDmxChan.Lock()
//	defer d.MtRwDmxChan.Unlock()
//	d.DmxChans = append(d.DmxChans, ch)
//
//	return nil
//}

func (d *DMX) runDemux(msg Message){

	d.MtRwDmxChan.RLock()
	for _, c := range d.DmxChans {
		c <- msg
	}
	d.MtRwDmxChan.RUnlock()

}


func (d *DMX) RunDMXOnce(){
	msg := <-d.inDmxChan
	d.runDemux(msg)
}

func (d *DMX)RunDMXOnceCloseAllChan(){
	msg := <- d.inDmxChan
	d.MtRwDmxChan.RLock()
	for _, c := range d.DmxChans {
		c <- msg
		close(c)
	}
	d.MtRwDmxChan.RUnlock()

}

func (d *DMX) clearChans(){
	//Взависимости от задач можно удалять только закрытые каналы
	d.DmxChans = make([]chan<- Message, 0)
}

func (d *DMX)RunDMXInf(){
	for {
		msg := <- d.inDmxChan
		go d.runDemux(msg)
	}
}

func (d *DMX) GetInChanDMX() chan <- Message {
	return d.inDmxChan
}

func (d *DMX)DmxClose(){
	close(d.inDmxChan)
}

