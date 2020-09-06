package main

import (
	"fmt"
	sort2 "sort"
	"test2/search"
)

func main() {
	var nums = []int{1,2,3,3,3,3,4,5}
	var nums2 = []int{0,1,3,4,5}
	var nums3 = []int{-1,1,3,4,5}
	if len(nums)<=0 || nums==nil {
		return
	}
	fmt.Println(getNumberofK(nums,3))
	if len(nums2)<=0 || nums2==nil {
		return
	}
	fmt.Println(getMissingNumber(nums2,len(nums2)))
	if len(nums3)<=0 || nums3==nil {
		return
	}
	if index,flag := getNumberSameAsIndex(nums3);flag{
		fmt.Println(index)
	}else {
		fmt.Println("不存在")
	}
}
func getFirstK(nums []int ,k,index int) (first int) {
	first=index
	for first>0&&nums[first-1]==k{
		first=search.SearchK(nums,0,first-1,k)
	}
	return first
}
func getLastK(nums []int,k,index int) (last int) {
	last=index
	for last<len(nums)-1&&nums[last+1]==k{
		last=search.SearchK(nums,last+1,len(nums)-1,k)
	}
	return last
}
//题目一：在排序数组中找到目的数字出现的次数
func getNumberofK(nums []int,k int) (count int) {
	index := search.SearchK(nums,0,len(nums)-1,k)
	if index ==-1 {
		return 0
	}
	first := getFirstK(nums,k,index)
	last := getLastK(nums,k,index)
	count=last-first+1
	return count
}
//题目二0—n-1中缺失的数字
func getMissingNumber(nums []int,n int) (index int) {
	if !sort2.IntsAreSorted(nums) {
		fmt.Println("未排序")
		return -1
	}
	index=-1
	start,end := 0,n
	mid := (start+end)/2
	for start<end{
		if nums[mid]==mid{
			start=mid+1
		}else if mid>0 {
			if nums[mid-1]==mid-1 {
				index=mid
				break
			}else {
				end=mid-1
			}
		}else {
			break
		}
		mid=(start+end)/2
	}
	if nums[0]==0 {
		return index
	}else {
		return 0
	}
}
//题目三：数组中数值和下标相等的元素
func getNumberSameAsIndex(nums []int) (index int,flag bool) {
	flag=false
	start,end :=0,len(nums)-1
	mid := (start+end)/2
	for start<=end{
		if mid==nums[mid] {
			flag=true
			index=mid
			break
		}else if mid>nums[mid] {
			start=mid+1
		}else {
			end=mid-1
		}
		mid=(start+end)/2
	}

	return index,flag
}

func getNumberofK2(nums []int,k int){
	var first=getFirstK2(nums,k)
	var last=getLastK2(nums,k)
	if first==-1||last==-1{
		fmt.Println(0)
		return
	}
	fmt.Println(last-first+1)
}
func getFirstK2(nums []int,k int) int {
	var i,j=0,len(nums)-1
	var mid int
	for i<=j{
		mid=(j+j)/2
		if nums[mid]==k{
			if (mid>0&&nums[mid-1]!=k)||mid==0{
				return mid
			}else {
				j=mid-1
			}
		}else if nums[mid]>k{
			j=mid-1
		}else {
			i=mid+1
		}
	}
	return -1
}
func getLastK2(nums []int,k int) int {
	var i,j=0,len(nums)-1
	var mid=0
	for i<=j{
		mid=(i+j)/2
		if nums[mid]==k{
			if (mid+1<len(nums)&&nums[mid+1]!=k)||mid==len(nums)-1{
				return mid
			}else {
				i=mid+1
			}
		}else if nums[mid]<k{
			i=mid+1
		}else {
			j=mid-1
		}
	}
	return -1
}

func getMissingNumber2(nums []int)  {
	var i,j=0,len(nums)-1
	var mid int
	for i<=j{
		mid=(i+j)/2
		if nums[mid]==mid{
			i=mid+1
		}else  {
			if (mid>0&&nums[mid-1]==mid-1)||mid==0{
				break
			}else {
				j=mid-1
			}
		}
	}
	if (mid==0&&nums[mid]!=mid)||(mid>0&&nums[mid-1]==mid-1&&nums[mid]!=mid){
		fmt.Println(mid)
		return
	}
	fmt.Println("not exist")
}
func getNumberSameAsIndex2(nums []int)  {
	var i,j=0,len(nums)-1
	var mid int
	for i<=j{
		mid=(i+j)/2
		if nums[mid]==mid{
			break
		}else if nums[mid]>mid{
			j=mid-1
		}else {
			i=mid+1
		}
	}
	if nums[mid]==mid{
		fmt.Println(mid)
		return
	}
	fmt.Println("not exist")
}