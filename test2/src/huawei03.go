package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int//N张机票
	fmt.Scan(&N)
	var str string
	var mmap=make(map[string]string)
	for i:=0;i<N;i++{
		fmt.Scan(&str)
		strs:=strings.Split(str,",")
		mmap[strs[0]+","+strs[1]]=strs[2]
	}
	//fmt.Println(mmap)
	var M int
	fmt.Scan(&M)
	for i:=0;i<M;i++{
		fmt.Scan(&str)
		strs:=strings.Split(str,",")
		if value,ok:=mmap[strs[0]+","+strs[1]];ok{
			delete(mmap,strs[0]+","+strs[1])
			mmap[strs[2]+","+strs[3]]=value
		}
	}
	//fmt.Println(mmap)
	var strs []string
	for key,value:=range mmap{
		strs=append(strs,key+","+value)
	}
	for i:=0;i<len(str);i++{
		fmt.Println(strs[i])
	}
}
//3
//CZ7132,A1,ZHANGSAN
//CZ7132,A2,ZHAOSI
//CZ7156,A2,WANGWU
//2
//CZ7132,A1,CZ7156,A2
//CZ7156,A2,CZ7156,A3
