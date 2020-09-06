package main

import "fmt"

func main() {
	var nums=[]int{1,1,2,2,0,1,1}
	fmt.Println(canJump(nums))
}
//func canJump(nums []int) bool {
//	if len(nums)==0 || nums==nil {
//		return false
//	}
//	if len(nums)==1{
//		return true
//	}
//	var length=len(nums)-1
//	for i:=length-1;i>=0;i--{
//		if nums[i]+i>=length&&canJump2(nums,i){
//			return true
//		}
//	}
//	return false
//}
//func canJump2(nums []int,i int) bool {
//	if i==0{
//		return true
//	}
//	for j:=i-1;j>=0;j--{
//		if nums[j]+j>=i&&canJump2(nums,j){
//			return true
//		}
//	}
//	return false
//}

func canJump(nums []int) bool {
	var index int
	for ;index<len(nums);{
		if index==len(nums)-1{
			return true
		}
		if nums[index]==0 {
			return false
		}
		index=canJump2(nums,index)
		fmt.Println("index=",index)
	}
	return true
}
func canJump2(nums []int,index int) int {
	var max int
	max=index+nums[index]
	if max>=len(nums)-1 {
		return len(nums)-1
	}
	var temp int
	var rs=index+1
	for i:=index+1;i<len(nums)&&i<=index+nums[index];i++{
		temp=nums[i]+i
		if temp>max {
			max=temp
			rs=i
		}
	}
	return rs
}


