package main

import (
	"fmt"
	"strings"
)

func main() {
	// your code
	//fmt.Println(ConvertFracts([][]int{{1, 2}, {1, 3}, {1, 4}}))
	fmt.Println(ConvertFracts([][]int{{69, 130}, {87, 1310}, {30, 40}}))
}


func ConvertFracts(a [][]int) string {
	// your code
	n, factor := nok(a)
	newDrop := make([][]int , len(a))
	for ind, drop := range a{
		newDrop[ind] = []int{drop[0]*(n/drop[1]), n}
	}

	newDrop = simplifyDrop(newDrop, factor)
	res := ""
	for _, drop := range newDrop{
		res += fmt.Sprintf("(%d,%d),", drop[0],drop[1])
	}
	return  strings.TrimRight(res, ",")
}

func simplifyDrop(drops[][]int, nok []int)[][]int{
	NewDrop := make([][]int, len(drops))
	copy(NewDrop, drops)
	//fmt.Println(drops)
	for _, num := range nok{
		 for i:=0;i<len(drops);i++{
			 if drops[i][0]%num == 0 && drops[i][1]%num == 0{
				 NewDrop[i] = []int{drops[i][0]/num, drops[i][1]/num}
			 }else{
				 copy(NewDrop, drops)
				 break
			 }

		 }
		copy(drops,NewDrop)
	}

	return drops
}




func nok(a [][] int)(int, []int){
	mn := make([][]int, len(a))
	for indx , num := range a{
		mn[indx] = factoring(num[1])
	}
	res := 1
	comFac := commonFactoring(mn)
	for _, cf := range comFac{
		res *= cf
	}
	return res, comFac
}



func factoring(n int) []int{
	nm := make([]int,0)
	for i:=2; i*i <= n;{
		if n%i == 0 {
			nm = append(nm, i)
			n /= i
			continue
		}
		i++
	}
	if n > 1{
		nm = append(nm, n)
	}
	return nm
}

func commonFactoring(factor [][]int)[]int{
	commonFact := factor[0]
	for _, f := range factor[1:]{

		for _, numCommon := range commonFact{
			 	for ind := 0 ;ind < len(f);ind++{
					 if f[ind] == numCommon{
						 f[ind] = -1
						 break
					 }
				}
		}
		for _, num := range f{
			if num != -1{
				commonFact = append(commonFact, num)
			}

		}
	}
	return commonFact
}