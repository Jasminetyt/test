package main

import (
	"fmt"
	"sort"
)

func main() {
	var candidates=[]int{10,1,2,7,6,2,2,2,1,5}
	var target=8
	fmt.Println(combinationSum2(candidates,target))
}

func combinationSum2(candidates []int, target int) [][]int {
	var (
		temp []int
		result [][]int
		sum=0
		index=0
	)
	sort.Ints(candidates)
	getCombinationSum2(candidates,target,sum,index,temp,&result)
	return result
}

func getCombinationSum2(candidates []int,target,sum,index int,temp []int,result *[][]int)  {
	if sum==target{
		var temp2=make([]int,len(temp))
		copy(temp2,temp)
		*result=append(*result,temp2)
		return
	}
	if sum>target{
		return
	}
	for i:=index;i<len(candidates);i++{
		if sum+candidates[i]>target{
			break
		}
		if len(temp)==0 ||(len(temp)>0&&candidates[i]>=temp[len(temp)-1]){
			temp=append(temp,candidates[i])
			getCombinationSum2(candidates,target,sum+candidates[i],i+1,temp,result)
			temp=temp[:len(temp)-1]
		}
		for i+1<len(candidates)&&candidates[i+1]==candidates[i]{
			i++
		}
	}
}
