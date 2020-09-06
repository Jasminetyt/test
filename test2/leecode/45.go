package main

import "fmt"

func main() {
	var nums=[]int{0,1}
	fmt.Println(jump(nums))
}
func jump(nums []int) int {
	if len(nums)==1{
		return 0
	}
	var index,count=0,0
	for ;index<len(nums);{
		index = jump2(nums,index,nums[index])
		fmt.Println("index=",index)
		count++
		if index==len(nums)-1{
			return count
		}
	}
	return count
}

func jump2 (nums []int,i,num int)(index int){
	var max,count=0,0
	var temp2 int
	if nums[i]+i>=len(nums)-1{
		return len(nums)-1
	}
	for j:=1;j<=num;j++{
		if i+j<len(nums){
			temp2 = i+j+nums[i+j]
		}else{
			count=len(nums)-1 //自己就可以跳到最后一个位置
			break
		}
		if temp2>max{
			max=temp2
			count=i+j
		}
	}
	return count
}
