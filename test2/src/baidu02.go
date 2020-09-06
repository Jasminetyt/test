package main

import "fmt"

func main() {
	var n,k int
	fmt.Scan(&n,&k)//输入 n和k
	var count int
	getcount1(n,k,&count)
	fmt.Println(count)
}
func getcount1(n,k int,count *int){ //count为全局变量
	//fmt.Println(n,k,*count)
	if n-k<=0 || (n-k)%2!=0{
		*count++
		return
	}
	temp:=(n-k)/2
	getcount1(temp,k,count)
	getcount1(temp+k,k,count)
}
