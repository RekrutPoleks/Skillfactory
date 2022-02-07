package sort

func bubbleSort(ar []int) {
	var flag bool
	//fmt.Println(flag)
	tmpsize := len(ar)
	for n :=0; len(ar) > n; n++{
		for i:=1; i < tmpsize;i++ {
			if ar[i-1] > ar[i]{
				ar[i-1], ar[i] = ar[i], ar[i-1]
				flag = true
			}
		}
		tmpsize--
		if flag != true{
			return
		}
		flag =false
	}

}

func selectionSort(ar []int) {
	for startIndex := 0; startIndex < len(ar)/2; startIndex++ {
		tmp := len(ar) - startIndex-1
		minIndex := startIndex
		maxIndex := len(ar) - startIndex-1
		for ind := startIndex + 1; ind <= tmp; ind++{
			if ar[ind] < ar[minIndex]{
				minIndex = ind
			}else if ar[ind] > ar[maxIndex]{
				maxIndex =  ind
			}
		}
		if minIndex != startIndex{
			ar[startIndex], ar[minIndex] = ar[minIndex], ar[startIndex]
		}
		if (maxIndex != tmp){
			ar[tmp], ar[maxIndex] = ar[maxIndex], ar[tmp]
		}
	}
}

func insertionSort(ar []int) {
	for indx := 1; indx < len(ar); indx++{
		//fmt.Println(indx)
		for i := indx; i > 0; i--{
			//fmt.Println(i)
			if ar[i-1] > ar[i]{
				ar[i], ar[i-1] = ar[i-1], ar[i]
			}else {
				break
			}
		}
	}
}

func mergeSort(ar []int) []int {
	// ваш код здесь
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Printf("recovered from panic: %v\n", r)
	//	}
	//}()

	if len(ar) == 1 {
		return ar
	} else {
		sl1, sl2 := ar[:len(ar)/2], ar[len(ar)/2:]
		return  merge(mergeSort(sl1), mergeSort(sl2))
	}
}
func merge(mergeAr1 []int, mergeAr2 []int)[]int{
	var j, k = 0, 0
	resMerge := make([]int, len(mergeAr1)+len(mergeAr2))
	//fmt.Println(mergeAr1, mergeAr2)
	for i := 0; i < len(resMerge); i++{
		switch {
		case len(mergeAr1) - 1 >= j && len(mergeAr2)-1 >= k:
			if mergeAr1[j] > mergeAr2[k] {
				resMerge[i] = mergeAr2[k]
				k++
			} else {
				resMerge[i] = mergeAr1[j]
				j++
			}
		case len(mergeAr1)-1 >= j:
			resMerge[i] = mergeAr1[j]
			j++
		default:
			//fmt.Println("default")
			resMerge[i] = mergeAr2[k]
			k++
		}
	}
	return resMerge
}

func quickSort(ar []int) {
	if len(ar) <= 2{
		return
	}else{
		left := 0
		ar[len(ar)-1], ar[len(ar)/2] = ar[len(ar)/2], ar[len(ar)-1]
		for ind := 0 ; ind < len(ar)-1; ind++{
			if ar[ind] <= ar[len(ar)-1]{
				ar[left], ar[ind] = ar[ind], ar[left]
				left++
			}
		}
		ar[len(ar)-1], ar[left] = ar[left], ar[len(ar)-1]
		quickSort(ar[:left])
		quickSort(ar[left:])
	}
}