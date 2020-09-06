package main

import "fmt"

func main()  {
	arr := selfDividingNumbers(1,22)
	for _,val := range  arr{
		fmt.Printf("val=%d  ",val)
	}
}
func selfDividingNumbers(left int, right int) []int {
	var result []int
	var num int
	for i := left; i<=right;i++ {
		num = i
		for num>0 {
			if num%10 !=0 && i%(num%10)==0{
				num=num/10
			}else {
				break;}
		}
		if num==0 {
			result=append(result, i)
		}
	}
	return result
}
