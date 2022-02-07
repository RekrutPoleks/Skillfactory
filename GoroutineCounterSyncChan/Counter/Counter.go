package Counter

import "fmt"

type Counter struct {
	chanSync chan bool
	chanExit chan bool
	count *int
	limitCount int
	goCount int
}

func NewCounter()*Counter{
	var countGo, limit int
	fmt.Print("Введите ограничение счетчика: ")
	fmt.Scanln(&limit)
	fmt.Print("Количество горутин: ")
	fmt.Scanln(&countGo)
	return &Counter{make(chan bool),make(chan bool), new(int),limit, countGo}
}


func(c *Counter)Start(){
	//c.chanSync <- true
	for ;c.goCount>0; c.goCount--{
		go c.add()
	}
	c.chanSync <- true
	<- c.chanExit
}

func (c *Counter)add(){
	for  {
		<- c.chanSync
		*(c.count) = *(c.count) + 1
		if c.isfull(){
			c.exit()
			<- c.chanSync
		}
		c.chanSync <- true
	}
}

func (c *Counter)isfull()bool{
	if c.limitCount == *c.count{
		return true
	} else {
		return false
	}
}

func (c *Counter)exit(){
	c.chanExit<-true
}

func (c *Counter)GetResult()int{
	return *c.count
}