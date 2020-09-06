package main

import "fmt"

func main() {
	var n,m int
	fmt.Scan(&n,&m)
	if n<1 || m<1{
		return
	}
	var last=0
	for i:=2;i<=n;i++{
		last=(last+m)%i
	}
	fmt.Println(last)
}

