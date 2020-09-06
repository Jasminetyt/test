package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var candidates=[]int{7,3,2}
	var target=18
	fmt.Println(combinationSum(candidates,target))
}
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var temp []int
	var sum=0
	getCombinationSum(candidates,target,&result,sum,temp)
	return result
}
//去重：先将数组进行排序，然后每次判断新加入temp的数字不能小于已有数字的最后一位
func getCombinationSum(candidates []int,target int,result *[][]int,sum int,temp []int)  {
	if sum==target{
		var temp2 =make([]int,len(temp))
		copy(temp2,temp)//将temp的值复制给temp2
		sort.Ints(temp2)
		flag :=true
		for i:=0;i<len(*result);i++{
			if reflect.DeepEqual(temp2,(*result)[i]){//判断两个切片是否相等，先判断类型，后判断值
				flag=false
			}
		}
		if flag {
			*result=append(*result,temp2)
		}
		return
	}
	if sum>target {
		return
	}
	for i:=0;i<len(candidates);i++{
		temp=append(temp,candidates[i])
		sum=sum+candidates[i]
		if sum>target {
			break
		}
		getCombinationSum(candidates,target,result,sum,temp)
		sum=sum-candidates[i]
		temp=temp[:len(temp)-1]
	}
}