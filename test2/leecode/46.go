package main

import "fmt"

func main() {
	var nums=[]int{1,2,3}
	fmt.Println(permute(nums))
}
func permute(nums []int) [][]int {
	if len(nums)==0 || nums==nil{
		return nil
	}
	var result [][]int
	permute2(nums,0,&result)
	return result
}

func permute2(nums []int,index int,result *[][]int){
	if index==len(nums) {
		temp2:= make([]int,len(nums))
		copy(temp2,nums)
		*result=append(*result,temp2)
		return
	}
	for i:=index;i<len(nums);i++{
		nums[index],nums[i]=nums[i],nums[index]
		permute2(nums,index+1,result)
		nums[index],nums[i]=nums[i],nums[index]
	}
}
