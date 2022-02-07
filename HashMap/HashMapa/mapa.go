package HashMapa

import "Skillfactory/HashMap/MyHashFunc"

type MyHashMapa struct{
	mymap [1000]string
}

func (h *MyHashMapa) Set(key, val string) {
	index := MyHashFunc.HashStr(key)
	h.mymap[index] = val
}

func (h *MyHashMapa) Get(key string) (value string, ok bool) {
	index := MyHashFunc.HashStr(key)
	if tmp := h.mymap[index]; tmp == ""{
		return "", false
	} else {
		return tmp, true
	}
}

func (h *MyHashMapa) Delete(key string) {
	index := MyHashFunc.HashStr(key)
	h.mymap[index] = ""
}

func NewHashMap()*MyHashMapa{
	return &MyHashMapa{}
}