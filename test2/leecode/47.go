package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums=[]int{0,1,0,0,9}
	fmt.Println(permuteUnique(nums))
	fmt.Println(len(permuteUnique(nums)))
}
func permuteUnique(nums []int) [][]int {
	if len(nums)==0 || nums==nil{
		return nil
	}
	sort.Ints(nums)
	var result [][]int
	permuteUnique2(nums,0,&result)
	return result
}
func permuteUnique2(nums []int,index int,result *[][]int)  {
	if index==len(nums){
		temp := make([]int,len(nums))
		copy(temp,nums)
		*result=append(*result,temp)
		return
	}
	var mnums=map[int]int{}
	for i:=index;i<len(nums);i++{
		if _,ok:=mnums[nums[i]];ok{
			continue
		}
		nums[i],nums[index]=nums[index],nums[i]
		permuteUnique2(nums,index+1,result)
		nums[i],nums[index]=nums[index],nums[i]
		mnums[nums[i]]=1
	}

}
