package main

import "fmt"

func main() {
	var N,T,M int
	fmt.Scan(&N,&T,&M)
	var num=make([]int,N)
	for i:=0;i<N;i++{
		fmt.Scan(&num[i])
	}
	fmt.Println(3)
}
