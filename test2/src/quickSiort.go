package main

import (
	"fmt"
)

func main() {
	nums := []int{3,4,1,6,2,7,4,3,3}
	QuickSort(nums,0,8)
	fmt.Println(nums)
}
func QuickSort(nums []int,left int,right int)  {
	if right>left {
		var temp =nums[left]
		var j=left+1
		for i :=left+1;i<=right;i++ { //交换次数过于多
			if nums[i]<temp {
				nums[i] ,nums[j]=nums[j],nums[i]
				j++
			}
		}
		j--;
		nums[left],nums[j]=nums[j],nums[left]
		QuickSort(nums,left,j-1)
		QuickSort(nums,j+1,right)
	}

}


