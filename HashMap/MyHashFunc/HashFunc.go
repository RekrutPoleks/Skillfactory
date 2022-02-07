package MyHashFunc

func HashInt64(key int)uint{
	return uint(key%1000)
}

func HashStr(key string)uint{
	var hassum uint
	for index, value := range key{
			hassum += uint(int32(index)*value)
	}
	return hassum%1000
}
