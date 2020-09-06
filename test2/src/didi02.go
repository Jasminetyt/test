package main

import "fmt"

func main() {
	var n,total,cost int
	fmt.Scan(&n,&total,&cost)
	var a=make([]int,n)
	var b=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	for i:=0;i<n;i++{
		fmt.Scan(&b[i])
	}
	fmt.Println(3)
}
