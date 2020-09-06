package main

import "fmt"

func main()  {
	var nums1  = []int{1,2}
	var nums2  = []int{3,4}
	fmt.Println(findMedianSortedArrays(nums1,nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
length := len(nums1)+len(nums2)
var i ,j,index=0,0,0
var mid float64
		for i<len(nums1)&&j<len(nums2)&&index-1<length/2 {
			if nums1[i]<nums2[j] {
				nums=append(nums, nums1[i])
				i++
				index++
			}else {
					nums=append(nums,nums2[j])
					j++
				    index++
			}
		}
	for i<len(nums1)&&index-1<length/2 {
		nums= append(nums, nums1[i])
		index++
		i++
	}
	for j<len(nums2)&&index-1<length/2 {
		nums= append(nums, nums2[j])
		index++
		j++
	}
	if length==1 {
		return float64(nums[index-1])
	}
	if length%2==0 {
		mid = float64(nums[index-1]+nums[index-2])/2
	}else {
		mid = float64(nums[index-1])
	}
	return mid
}
