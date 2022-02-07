package CasheMap

import (
	"fmt"
	"sync"
	"time"
)

var _ Cache = InMemoryCache{} // это трюк для проверки типа: до тех пор пока InMemoryCache не будет реализовывать интерфейс Cache, программа не запустится

type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type InMemoryCache struct {
	//Code here
	duration time.Duration
	mutex sync.Mutex
	cashe map[string]*CacheEntry
}

func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		// Code here
		expireIn,
				sync.Mutex{},
		make(map[string]*CacheEntry),
	}
}


func (inMC *InMemoryCache) Set(key string, value interface{}){
	inMC.mutex.Lock()
	inMC.cashe[key] = &CacheEntry{
		time.Now(),
		value,
	}
	inMC.mutex.Unlock()
}
func (inMC *InMemoryCache) Get(key string)interface{}{
	incashe, ok := inMC.cashe[key]
	if !ok && time.Since(incashe.settledAt) < inMC.duration{
		return incashe
	}else{
		return fmt.Errorf("Alarm")
	}

}