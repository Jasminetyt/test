package main

import (
	"fmt"
	s1 "sort"
	)

func main()  {
	var n int
	fmt.Scan(&n) //人数
	var Anum=make([]int,n)
	var Bnum=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&Anum[i])
	}
	for i:=0;i<n;i++{
		fmt.Scan(&Bnum[i])
	}
	var i,j,first int
	var last=n-1
	var mapB=make(map[int][]int)
	for i=0;i<n;i++{
		mapB[Bnum[i]]=append(mapB[Bnum[i]],i)
	}
	s1.Ints(Anum)
	s1.Ints(Bnum)
	var answer=make([]int,n)
	i,j=0,0
	for i<n&&j<n{
		if Anum[i]>Bnum[j]{
			answer[first]=Anum[i]
			first++
			i++
			j++
		}else {
				answer[last]=Anum[i]
				last--
				i++
		}
	}
	j=0
	var result=make([]int,n)
	for i=0;i<len(answer);i++{
		temp:=mapB[Bnum[i]][j]
		for temp==-1{
			j++
			temp=mapB[Bnum[i]][j]
		}
		mapB[Bnum[i]][j]=-1
		j=0
		result[temp]=answer[i]
	}
	var sum int
	for i=0;i<n;i++{
		if answer[i]>Bnum[i] {
			sum=sum+1
		}else if answer[i]<Bnum[i]{
			sum=sum-1
		}
	}
	fmt.Println(sum)
	//for ;i<n&&j<n;{
	//	if Anum[i]>Bnum[j] {
	//		i++
	//		j++
	//		result=result+1
	//	}else {
	//		i++
	//	}
	//}
	//result=result-i+j
	////fmt.Println(result)
	//max:=result
	//i,j,result=0,0,0
	//for i<n&&j<n{
	//	if Anum[i]>=Bnum[j]{
	//		if Anum[i]>Bnum[j] {
	//			result=result+1
	//		}
	//		i++
	//		j++
	//	}else {
	//			i++
	//	}
	//}
	//result=result-i+j
	//if result>max {
	//	max=result
	//}
	//fmt.Println(max)
}