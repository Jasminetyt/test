package main

import "fmt"

func main() {
	var nums=[]int{2}
	var val=3
	length := removeElement(nums,val)
	fmt.Println(length,nums)
}
func removeElement(nums []int, val int) int {
	if nums==nil || len(nums)==0{
		return 0
	}
	var i,j =0,len(nums)-1
	for i<=j{
		for i<len(nums)&&nums[i]!=val{
			i++
		}
		for j>=0&&nums[j]==val{
			j--
		}
		fmt.Println(i,j)
		if i<j{
			nums[i],nums[j]=nums[j],nums[i]
		}
	}
	return i
}