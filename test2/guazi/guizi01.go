package main

import "fmt"

func main() {
	var nums=[]int{1,2,3,4,5,6}
	index,flag:=search(nums,4)
	if flag{
		fmt.Println(nums[index],"exist")
	}else {
		fmt.Println("not exist")
	}
}
func search(nums []int,k int) (int,bool) {
	var start,end=0,len(nums)-1
	var mid int
	for start<=end{
		mid=(start+end)/2
		if nums[mid]==k{
			break
		}else if nums[mid]<k {
			start=mid+1
		}else {
			end=mid-1
		}
	}
	var flag bool
	if nums[mid]==k{
		flag=true
	}
	return mid,flag
}
