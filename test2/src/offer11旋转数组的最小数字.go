package main

import "fmt"

func main() {
	nums := []int{3,3,4,5,2,3}
	fmt.Println(revolve(nums,0,len(nums)-1))
	fmt.Println(revolve2(nums))
}
//递归
func revolve(nums []int,left,right int) int {
	if right==-1 {
		return -1
	}
	if left==right-1 {
		return nums[right]
	}
	mid := (left+right)/2
	fmt.Printf("left=%d,right=%d,mid=%d \n",left,right,mid)
	if nums[left]<nums[right] {
		return nums[left]
	}
	if nums[mid]>=nums[left]&&nums[mid]>nums[right]{
		left=mid
	}else if nums[mid]<=right&&nums[mid]<nums[left] {
		right=mid
	}else if nums[mid]==nums[right]&&nums[mid]==nums[left] {
		fmt.Println("left==mid==right")
		return selectt(nums,left,right)
	}
	return revolve(nums ,left,right)
}

func selectt(nums []int ,left,right int) int {
	for i:=left+1;i<=right-1;i++ {
		if nums[i]<nums[i-1]&&nums[i]<nums[i+1] {
			return nums[i]
		}
	}
	if nums[left]<nums[right]{
		return nums[left]
	}else {
		return nums[right]
	}
}
//循环
func revolve2(nums []int) int {
	if len(nums)==0 {
		return -1
	}
	left ,right := 0,len(nums)-1
	if nums[left]<nums[right] {
		return nums[left]
	}
	for left<right-1  {
		mid := (left+right)
		if nums[mid]==nums[left]&&nums[mid]==nums[right] {
			return selectt(nums,left,right)
		}else if(nums[left]>=nums[mid]){
			right=mid
		}else {
			left=mid
		}
	}
	return nums[right]
}

func resolvemin(nums []int) int {
	start,end:=0,len(nums)-1
	var mid =start
	for nums[start]>=nums[end]{
		if end-start==1{
			mid=end
			break
		}
		mid = (start+end)/2
		if nums[start]==nums[end]&&nums[mid]==nums[start]{
			return minorder(nums,start,end)
		}
		if nums[mid]<nums[end]{
			end=mid
		}else {
			start=mid
		}
	}
	return nums[mid]
}
func minorder(nums []int,start,end int) int {
	var result=nums[start]
	for i:=start+1;i<=end;i++{
		if result>nums[i]{
			result=nums[i]
		}
	}
	return result
}