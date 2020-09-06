package main

import "fmt"

func main() {
	var N,K int
	fmt.Scan(&N,&K)
	var mmap=make(map[int]int)
	var key,value int
	for i:=0;i<N;i++{
		fmt.Scan(&key,&value)
		mmap[key]=value
	}
	var result=make([]int,K)
	var temp int
	for i:=0;i<K;i++{
		fmt.Scan(&temp)
		result[i]=mmap[temp]
	}
	for i:=0;i<K;i++{
		fmt.Println(result[i])
	}
}
