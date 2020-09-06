package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
	var result=make([][]int,T)
	for i:=0;i<T;i++{
		result[i]=make([]int,2)
	}
	var n,sum int
	for i:=0;i<T;i++{
		sum=0
		fmt.Scan(&n)
		var num=make([]int,n)
		for i:=0;i<n;i++{
			fmt.Scan(&num[i])
			sum=sum+num[i]
		}
		sort.Ints(num)

}
	}
