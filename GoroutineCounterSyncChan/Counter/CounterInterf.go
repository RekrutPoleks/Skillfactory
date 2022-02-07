package Counter

type CounterInterface interface{
	Start()
	add()
	isfull()bool
	exit()
	GetResult()int
}
