package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n,m int
	fmt.Scan(&n,&m)
	var count int64
	var temp []int
	var mnum=make(map[string]int)
	count1(n,m,&count,temp,mnum)
	fmt.Println(count)
}
func count1(n,m int,count *int64,temp []int,mnum map[string]int)  {
	if len(temp)==m{
		var temp2=make([]int,m)
		copy(temp2,temp)
		sort.Ints(temp2)
		var str string
		for i:=0;i<len(temp2);i++{
			str=str+strconv.Itoa(temp2[i])
		}
		if _,ok:=mnum[str];!ok{
			//fmt.Println(str)
			*count++
			if *count>1000000007{
				*count=*count-1000000007
			}
			mnum[str]=1
		}
		return
	}
	for i:=1;i<=n;i++{
		temp=append(temp,i)
		count1(n,m,count,temp,mnum)
		temp=temp[0:len(temp)-1]
	}
}
