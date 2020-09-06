package main

import "fmt"

func main() {
	var t int
	fmt.Println(&t)
	var result=make([]string,t)
	var n int
	for i:=0;i<t;i++{
		fmt.Scan(&n)
		var num=make([]int,n)
		for j:=0;j<n;j++{
			fmt.Scan(&num[j])
		}
		result[i]=panduan(num)
	}
	for i:=0;i<t;i++{
		fmt.Println(result[i])
	}
}
func panduan(num []int) string {
	return "YES"
}